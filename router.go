package mano

import (
	"fmt"
	"reflect"
	"regexp"
)

type Handler interface{}
type HandlerFunc func(Context) interface{}
type ServiceFunc func(Context) reflect.Value
type HttpMethod string

const (
	GET  HttpMethod = "GET"
	POST HttpMethod = "POST"
)

type Router struct {
	method      HttpMethod
	pattern     string
	middlewares []string
	handler     Handler
	handle      HandlerFunc
}

func buildHandler(app *Application, handler Handler) (handle HandlerFunc) {
	handlerType := reflect.TypeOf(handler)
	switch handlerType.Kind() {
	case reflect.Func:
		params := make([]ServiceFunc, handlerType.NumIn())
		for i := 0; i < handlerType.NumIn(); i++ {
			//ptyp := handlerType.In(i)
			typeName := handlerType.In(i).String()
			fn, ok := app.paramterMap[typeName]
			if !ok {
				switch typeName {
				case "*mano.Application":
					param := reflect.ValueOf(app)
					fn = func(ctx Context) reflect.Value {
						return param
					}
				case "mano.Context":
					fn = func(ctx Context) reflect.Value {
						return reflect.ValueOf(ctx)
					}
				case "*mano.Request":
					fn = func(ctx Context) reflect.Value {
						return reflect.ValueOf(ctx.Request())
					}
				case "*mano.Response":
					fn = func(ctx Context) reflect.Value {
						return reflect.ValueOf(ctx.Response())
					}
				case "*http.Request":
					fn = func(ctx Context) reflect.Value {
						return reflect.ValueOf(ctx.Request().RawRequest())
					}
				case "http.ResponseWriter":
					fn = func(ctx Context) reflect.Value {
						return reflect.ValueOf(ctx.Response().Writer())
					}
				default:
					for _, val := range app.attrs {
						vt := reflect.TypeOf(val)
						if vt.String() == typeName {
							param := reflect.ValueOf(val)
							fn = func(ctx Context) reflect.Value {
								return param
							}
						}
					}
					break
				}

				if fn == nil {
					panic(fmt.Sprintf("参数不支持或未定义:%v", typeName))
				}
				app.paramterMap[typeName] = fn

			}
			params[i] = fn
		}

		handle = func(ctx Context) interface{} {
			method := reflect.ValueOf(handler)

			arr := make([]reflect.Value, len(params))
			for i := 0; i < len(params); i++ {
				arr[i] = params[i](ctx)
			}
			returns := method.Call(arr)
			if len(returns) > 0 {
				return returns[0].Interface()
			}
			return nil
		}

	}
	return
}

var patternRegexp = regexp.MustCompile(`\{\s*(\w+)\s*(\??)\s*\}`)

func compilePattern(pattern string) string {

	var index []int
	var matched []string
	for {
		index = patternRegexp.FindStringIndex(pattern)
		if len(index) != 2 {
			break
		}
		matched = patternRegexp.FindStringSubmatch(pattern)

		if len(matched) != 3 {
			break
		}
		partten := "(?P<" + matched[1] + ">\\w+)"
		if matched[2] == "?" {
			partten += "?"
			//TODO:如果后面有空格的话将会判断不到
			//t.Fatalf("****：%v----%v", index[1]+1, len(pattern))
			if len(pattern) == index[1]+1 && pattern[index[1]] == '/' {
				index[1] = index[1] + 1
				partten += "/?"
			}
		}
		pattern = pattern[:index[0]] + partten + pattern[index[1]:]
		// t.Fatalf("****：%v", len(matched))
		//break
	}
	return pattern
}
