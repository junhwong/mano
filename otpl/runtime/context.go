package runtime

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/junhwong/mano/otpl/common"
)

type ContextData map[string]interface{}

func (data ContextData) Set(key string, value interface{}) interface{} {
	data[key] = value
	return value
}

func (data ContextData) Get(key string) (value interface{}, ok bool) {
	value, ok = data[key]
	return
}

func (data ContextData) Del(key string) (value interface{}, ok bool) {
	value, ok = data[key]
	delete(data, key)
	return
}

func (data ContextData) Len() int {
	return len(data)
}

type contextScope struct {
	data   ContextData
	stack  common.Stack
	parent *contextScope
}

func newContextScope(data ContextData, parent *contextScope) *contextScope {
	scope := &contextScope{
		stack:  *new(common.Stack),
		data:   make(ContextData, 0),
		parent: parent,
	}
	if parent != nil {
		for key, val := range parent.data {
			scope.data.Set(key, val)
		}
	}

	if data != nil {
		for key, val := range data {
			scope.data.Set(key, val)
		}
	}
	return scope
}

type context struct {
	currentScope *contextScope
	path         string
	out          io.Writer
	runtime      *Runtime
	loaders      map[string]common.Loader
	scopes       []contextScope // common.Stack
}

func (ctx *context) IsStrict() bool {
	return true //ctx.runtime.strictMode
}

func (ctx *context) Pop() interface{} {
	if v, ok := ctx.currentScope.stack.Pop(); ok {
		return v
	}
	return nil
}

func (ctx *context) Push(val interface{}) interface{} {
	return ctx.currentScope.stack.Push(val)
}

func (ctx *context) Var(name string, val ...interface{}) (interface{}, bool) {
	name = strings.ToLower(name) //忽略大小写

	if val != nil && len(val) > 0 {
		return ctx.currentScope.data.Set(name, val[0]), true
	}
	//  else if v, ok := ctx.currentScope.data.Get(name); ok {

	// 	return v
	// }
	return ctx.currentScope.data.Get(name)
}

func ToString(args ...interface{}) string {
	result := ""
	for _, arg := range args {
		if arg != nil {
			result += fmt.Sprintf("%v", arg)
		}
	}
	return result
}

func (ctx *context) Print(obj interface{}, escape bool) error {
	if obj == nil {
		return nil
	}
	_, err := ctx.out.Write([]byte(ToString(obj)))
	if err != nil {
		return err
	}
	return nil
	// if s, ok := obj.(string); ok {
	// 	_, err := ctx.out.Write([]byte(s))
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// fmt.Printf("XXXX:%v\n", reflect.TypeOf(obj))
	// //TODO:完善
	// return nil
}

func (ctx *context) Scope() error {
	//ctx.scopes.Push(ctx.currentScope)
	ctx.scopes = append(ctx.scopes, *ctx.currentScope)

	ctx.currentScope = newContextScope(nil, ctx.currentScope)

	return nil
}

func (ctx *context) Unscope() error {
	if len(ctx.scopes) == 0 {
		return errors.New("Invalid opertion,scope stack is empty")
	}
	// ctx.currentScope.free()
	ctx.currentScope = &ctx.scopes[len(ctx.scopes)-1]
	ctx.scopes = ctx.scopes[:len(ctx.scopes)-1]
	// if val, ok := ctx.scopes.Pop(); ok {
	// 	if scope, ok2 := val.(contextScope); ok2 {
	// 		ctx.currentScope = &scope
	// 		return
	// 	}
	// 	vv := contextScope(val)
	// 	panic(fmt.Errorf("Missmatch context scope:%v\n", val))
	// }
	if ctx.currentScope != nil {
		return nil
	}
	return errors.New("Invalid opertion,scope stack is empty2")
}

func (ctx *context) Free() {
	for _, loader := range ctx.loaders {
		loader.Close()
	}
	ctx.loaders = nil
	ctx.scopes = nil
	ctx.currentScope = nil
}

func (ctx *context) Load(src, ref string) (common.Loader, error) {

	path := canonicalName(src, ref)
	//panic(path)
	key := path + ref
	if loader, ok := ctx.loaders[key]; ok {
		return loader, nil
	}

	filename := ctx.runtime.ilpath + "\\" + md5Sign(path) + ".otil"
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	//defer file.Close()
	// file.s

	loader, err := Open(file)
	if err != nil {
		return nil, err
	}
	ctx.loaders[key] = loader
	return loader, nil
}

func (ctx *context) Exec(loader common.Loader, ptr common.Ptr) (err error) {
	var code common.Opcode
	// var code2 common.Opcode
	for ptr > common.ZeroPtr {

		code, err = loader.Load(ptr)
		if err != nil {
			//panic(fmt.Sprintf("%v", reflect.TypeOf(code)))
			return
		}
		// code = code2
		//fmt.Printf("%v\n", reflect.TypeOf(code))
		ptr, err = code.Exec(ctx)
		if err != nil {
			//TODO: 处理异常，如加上 行号

			return
		}

	}
	return
}

func (ctx *context) TemplateFunc(name string, fn ...common.TemplateFunc) common.TemplateFunc {
	f, ok := funcs[name]
	if ok {
		return f
	}
	return nil
}

var funcs = make(map[string]common.TemplateFunc, 0)

func init() {

	funcs["len"] = func(params ...interface{}) (result interface{}, err error) {
		//doto:错误解决
		if len(params) == 0 || params[0] == nil {
			result = 0
		} else {
			result = reflect.ValueOf(params[0]).Len()
		}

		return
	}

}
