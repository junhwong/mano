package common

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Number interface {
	Int() int64
	Float() float64
	IsFloat() bool
}

type number struct {
	origin  interface{}
	isFloat bool
	i64     int64
	f64     float64
}

func (num *number) Int() int64 {
	if num.isFloat {
		return int64(num.f64)
	}
	return num.i64
}

func (num *number) Float() float64 {
	if !num.isFloat {
		return float64(num.i64)
	}
	return num.f64
}
func (num *number) IsFloat() bool {
	return num.isFloat
}

func ToNumber(v interface{}) (r Number, err error) {
	num := &number{
		origin:  v,
		isFloat: false,
	}
	r = num

	// vv := reflect.ValueOf(v)

	// kk := vv.Convert(reflect.TypeOf(""))

	// panic(fmt.Sprintf("xxxxxxxxxxxxxx:%v", kk.Interface()))

	switch n := v.(type) {
	case int:
		num.i64 = int64(n)
	case int8:
		num.i64 = int64(n)
	case int16:
		num.i64 = int64(n)
	case int32:
		num.i64 = int64(n)
	case int64:
		num.i64 = int64(n)
	case uint:
		num.i64 = int64(n)
	case uint8:
		num.i64 = int64(n)
	case uint16:
		num.i64 = int64(n)
	case uint32:
		num.i64 = int64(n)
	case uint64:
		num.i64 = int64(n)
	case float32:
		num.f64 = float64(n)
		num.isFloat = true
	case float64:
		num.f64 = float64(n)
		num.isFloat = true
	case string:
		s := string(n)
		if strings.IndexAny(s, ".") > 0 {
			num.f64, err = strconv.ParseFloat(s, 64)
			num.isFloat = true
		} else {
			num.i64, err = strconv.ParseInt(s, 10, 64)
		}
	default:
		err = fmt.Errorf("cannot convert to number(float64):%v(%v)", v, reflect.TypeOf(v))
	}

	// if i, ok := v.(int64); ok {
	// 	num.i64 = i
	// 	num.isFloat = false
	// } else if i, ok := v.(float64); ok {
	// 	num.f64 = i
	// 	num.isFloat = true
	// } else if s, ok := v.(string); ok {
	// 	num.i64, err = strconv.ParseInt(s, 10, 64)
	// 	num.isFloat = false
	// 	if err != nil {
	// 		num.f64, err = strconv.ParseFloat(s, 64)
	// 		num.isFloat = true
	// 	}
	// } else {
	// 	err = fmt.Errorf("cannot convert to number(float64):%v(%v)", v, reflect.TypeOf(v))
	// }

	return
}
