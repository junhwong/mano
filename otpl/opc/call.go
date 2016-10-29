package opc

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/junhwong/mano/otpl/common"
)

type opCall struct {
	opBase
	// parameterCount int
}
type Object interface{}

func (op *opCall) Load() (err error) {
	return nil
}

func (op *opCall) Exec(ctx common.Context) (ptr common.Ptr, err error) {
	ptr = op.ptr + 1
	parameterCount := int(op.flag)
	var fn common.TemplateFunc
	var obj Object
	method := ctx.Pop()

	if s, ok := method.(string); ok { //内置方法？
		fn = ctx.TemplateFunc(s)
		if fn == nil {
			if ctx.IsStrict() {
				err = fmt.Errorf("Undefined function:%s", s)
				return
			}
			//消耗参数
			for i := 0; i < parameterCount; i++ {
				ctx.Pop()
			}
			ctx.Push(nil)
			return
		}
	} else if method, ok := method.(*reflect.Value); ok {
		obj = ctx.Pop()
		fn = func(v ...interface{}) (interface{}, error) {
			//panic(len(v))
			//TODO:错误处理
			params := make([]reflect.Value, len(v)+1)
			for i := 1; i < len(params); i++ {
				params[i] = reflect.ValueOf(v[i-1])
			}
			params[0] = reflect.ValueOf(obj)
			returns := method.Call(params)
			if len(returns) > 0 {
				return returns[0].Interface(), nil //只返回第一个值
			}
			return nil, nil
		}
	}

	params := make([]interface{}, parameterCount)
	for i := parameterCount - 1; i >= 0; i-- {
		params[i] = ctx.Pop()
	}
	// params[0] = methodObj
	//panic(params[0])
	if fn == nil {
		err = handErr(true, ctx, errors.New("cannot call null template function"))
		return
	}

	r, err := fn(params...)
	if handErr(err != nil, ctx, err) == nil {
		ctx.Push(r)
	}

	return op.ptr + 1, nil
}

func init() {
	common.RegisterOpcode(0x07, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opCall{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
