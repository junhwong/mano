package mano

import (
	"fmt"
	"reflect"
	"regexp"
	"sync"

	"github.com/junhwong/mano/utils"
)

//Application ...
type Application struct {
	muLock            *sync.Mutex
	isDefault         bool
	middlewares       map[string]Middleware
	globalMiddlewares map[string]Middleware
	paramterMap       map[string]ServiceFunc
	handlers          []HTTPHandler
	attrs             utils.Attribute
	viewEngine        ViewEngine
	routeTable        *RouteTable
}

//Use 设置或注册全局变量、中间件、服务等
func (app *Application) Use(registrations ...interface{}) {
	app.muLock.Lock()
	defer app.muLock.Unlock()

	var err error
	c := len(registrations)
	if c == 1 && registrations[0] != nil {

		switch any := registrations[0].(type) {
		case ViewEngine:
			app.viewEngine = ViewEngine(any)
			return
		case HTTPHandler:
			handler := HTTPHandler(any)
			handler.Init(app)
			app.handlers = append(app.handlers, handler)
			return
		case *RouterGroup:
			group := (*RouterGroup)(any)
			for _, r := range group.items {
				pattern := regexp.MustCompile(compilePattern(r.pattern))
				handler := buildHandler(app, r.handler)
				app.routeTable.Register(pattern, handler, 0, false)
			}
			return
		default:
			err = fmt.Errorf("unsupport :%v", reflect.TypeOf(any).String())
			break
		}
	} else if c == 2 && registrations[0] != nil {
		if key, ok := registrations[0].(string); ok {
			if value, ok := registrations[1].(string); ok {
				app.attrs.SetProperty(key, value)
			} else {
				app.attrs.Set(key, registrations[1])
			}
			return
		} else {
			err = fmt.Errorf("invalid parameters :%v", registrations)
		}
	} else {
		err = fmt.Errorf("invalid parameters :%v", registrations)
	}

	panic(err)
}

func (app *Application) Items() utils.Attribute {
	return app.attrs
}

func (app *Application) Attr(key string, defaultValue ...string) string {
	if val := app.attrs.GetString(key); val != "" {
		return val
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}
