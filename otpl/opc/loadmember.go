package opc

import (
	"reflect"
	"sort"
	"strings"

	"github.com/junhwong/mano/otpl/common"
	"github.com/junhwong/mano/utils"
)

type opLoadMember struct {
	opBase
	getmethod bool
}

func (op *opLoadMember) Load() (err error) {
	op.getmethod, err = op.loader.ReadBool()
	return
}

func getFiled(obj *reflect.Value, name string) *reflect.Value {
	if obj.Kind() == reflect.Ptr {
		elem := obj.Elem()
		obj = &elem
	}

	member := obj.FieldByName(name)
	if member.IsValid() {
		return &member
	}

	member = obj.FieldByNameFunc(func(s string) bool {
		return strings.EqualFold(s, name)
	})

	if member.IsValid() {
		return &member
	}

	return nil
}

func getMethod(obj *reflect.Value, name string, paramsCount int) *reflect.Value {
	if obj.CanAddr() {
		addr := obj.Addr()
		obj = &addr
	}
	member := obj.MethodByName(name)
	if member.IsValid() && !member.IsNil() {
		return &member
	}
	typ := obj.Type()
	//typ.Elem()
	//panic(obj.NumMethod())
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		// fmt.Printf("xxxx %s=%s:%v\n", method.Name, name, strings.EqualFold(method.Name, name))
		if strings.EqualFold(method.Name, name) {
			return &method.Func
		}
	}

	return nil
}

func (op *opLoadMember) Exec(ctx common.Context) (ptr common.Ptr, err error) {
	err = nil
	ptr = op.ptr + 1

	instance := ctx.Pop()
	err = handErr(instance == nil, ctx, "could not load an member or null object dereference")
	if err != nil {
		return
	}
	err = handErr(utils.IsPrimitive(instance), ctx, "could not load with primitive object")
	if err != nil {
		return
	}

	parameterCount := int(op.flag)
	params := make([]interface{}, parameterCount)
	for i := 0; i < parameterCount; i++ {
		params[i] = ctx.Pop()
	}

	//对参数位置进行纠正
	si := utils.SortInterface(parameterCount, func(i, j int) bool {
		return i < j
	}, func(i, j int) {
		params[i], params[j] = params[j], params[i]
	})

	sort.Sort(sort.Reverse(si))

	obj := reflect.ValueOf(instance) //.Elem()

	if obj.Kind() == reflect.Map {
		keys := obj.MapKeys()
		var key *reflect.Value
		key, err = converWithKind(keys[0].Kind(), params[0])
		err = handErr(err != nil, ctx, err)
		if err != nil {
			return
		}

		ctx.Push(obj.MapIndex(*key).Interface())
	} else if obj.Kind() == reflect.Array {
		var n common.Number
		n, err = common.ToNumber(params[0])
		err = handErr(err != nil, ctx, err)
		if err != nil {
			return
		}
		size := obj.Len()
		index := int(n.Int())
		err = handErr(index >= size, ctx, "Array index greater than size:%d/%d", index, size)
		if err != nil {
			return
		}
		ctx.Push(obj.Index(index).Interface())
	} else if s, ok := params[0].(string); ok {

		member := getFiled(&obj, s)

		if member == nil {
			member = getMethod(&obj, s, parameterCount)
		}

		err = handErr(member == nil, ctx, "Member not found:%s", s)
		if err != nil {
			return
		}

		if op.getmethod { //
			ctx.Push(instance)
			// panic(instance)
		}

		//TODO:是取值还是函数？
		ctx.Push(member)
	}
	err = handErr(err != nil, ctx, "Unsupport object type(%s) or invalid param (%v)", obj.Type().String(), params[0])
	return
}

func init() {
	common.RegisterOpcode(0x0A, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opLoadMember{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
