package opc

import (
	"reflect"

	"github.com/junhwong/mano/otpl/common"
)

type Iterator interface {
	HasNext() bool
	Next() interface{}
	SetVariables(keyName, valName string)
}

type mapIterator struct {
	instance reflect.Value
	keys     []reflect.Value
	index    int
	size     int
	ctx      common.Context
}

func (iter *mapIterator) currentKey() interface{} {
	return iter.keys[iter.index].Interface()
}

func (iter *mapIterator) currentValue() interface{} {
	val := iter.instance.MapIndex(iter.keys[iter.index])
	if val.IsValid() && !val.IsNil() {
		return val.Interface()
	}
	return nil
}

func (iter *mapIterator) HasNext() bool {
	return iter.index+1 < iter.size
}

func (iter *mapIterator) Next() interface{} {
	iter.index++
	return iter.currentValue()
}

func (iter *mapIterator) SetVariables(keyName, valName string) {
	if keyName != "" {
		iter.ctx.Var(keyName, iter.currentKey())
	}
	if valName != "" {
		iter.ctx.Var(valName, iter.currentValue())
	}
}

type arrayIterator struct {
	instance reflect.Value
	index    int
	size     int
	ctx      common.Context
}

func (iter *arrayIterator) currentKey() interface{} {
	return iter.index
}

func (iter *arrayIterator) currentValue() interface{} {
	if iter.index < 0 || iter.index > iter.size-1 {
		return nil
	}
	//panic(iter.index)
	//return nil
	val := iter.instance.Index(iter.index)
	if val.IsValid() {
		return val.Interface()
	}
	return nil
}

func (iter *arrayIterator) HasNext() bool {
	//fmt.Printf("HasNext:%v\n", iter.index+1 < iter.size)
	return iter.index+1 < iter.size
}

func (iter *arrayIterator) Next() interface{} {
	iter.index++
	//fmt.Printf("Next:%v\n", iter.index)
	return iter.currentValue()
}

func (iter *arrayIterator) SetVariables(keyName, valName string) {
	if keyName != "" {
		iter.ctx.Var(keyName, iter.currentKey())
	}
	if valName != "" {
		iter.ctx.Var(valName, iter.currentValue())
	}
}

type nullIterator struct{}

func (iter *nullIterator) currentKey() interface{} {
	return nil
}

func (iter *nullIterator) currentValue() interface{} {
	return nil
}

func (iter *nullIterator) HasNext() bool {
	return false
}

func (iter *nullIterator) Next() interface{} {
	return nil
}

func (iter *nullIterator) SetVariables(keyName, valName string) {}

func newIterator(obj interface{}, ctx common.Context) Iterator {

	if iter, ok := obj.(Iterator); ok {
		return iter
	}

	var val reflect.Value
	if v, ok := obj.(reflect.Value); ok {
		val = v
	} else {
		val = reflect.ValueOf(obj)
	}

	switch val.Kind() {
	case reflect.Map:

		iter := &mapIterator{
			instance: val,
			keys:     val.MapKeys(),
			index:    -1,
			size:     val.Len(),
			ctx:      ctx,
		}
		return iter
	case reflect.Array, reflect.Slice:

		iter := &arrayIterator{
			instance: val,
			index:    -1,
			size:     val.Len(),
			ctx:      ctx,
		}
		return iter
	}
	panic(val.Kind())
	//unsupport other case
	return &nullIterator{}
}

type rangeIterator struct {
	instance reflect.Value
	ctx      common.Context
	start    int64
	stop     int64
	step     int64
}

func (iter *rangeIterator) currentKey() interface{} {
	return iter.start
}

func (iter *rangeIterator) currentValue() interface{} {
	return iter.start
}

func (iter *rangeIterator) HasNext() bool {
	return iter.start+1 < iter.stop
}

func (iter *rangeIterator) Next() interface{} {
	iter.start++
	return iter.currentValue()
}

func (iter *rangeIterator) SetVariables(keyName, valName string) {
	if keyName != "" {
		iter.ctx.Var(keyName, iter.currentKey())
	}
	if valName != "" {
		iter.ctx.Var(valName, iter.currentValue())
	}
}

func newRangeIterator(start, stop, step int64, ctx common.Context) *rangeIterator {
	return &rangeIterator{
		start: start,
		step:  step,
		stop:  stop,
		ctx:   ctx,
	}
}
