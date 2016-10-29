package opc

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/junhwong/mano/otpl/common"
	"github.com/junhwong/mano/utils"
)

type opBase struct {
	codeType   byte
	loader     common.Loader
	ptr        common.Ptr
	lineNumber common.LineNo
	flag       byte
}

// func (b *opBase) loadLineNumber() error {
// 	i, err := b.loader.ReadInt()
// 	if err == nil {
// 		b.lineNumber = i
// 	}
// 	return err
// }

func handErr(b bool, ctx common.Context, err interface{}, params ...interface{}) error {
	if !b {
		return nil
	}
	if !ctx.IsStrict() {
		ctx.Push(nil)
		return nil
	}
	if e, ok := err.(error); ok {
		return e
	} else if s, ok := err.(string); ok {
		if len(params) == 0 {
			return errors.New(s)
		}
		return fmt.Errorf(s, params...)
	}
	panic("Invalid paramter: err")
}

func converWithKind(kind reflect.Kind, v interface{}) (val *reflect.Value, err error) {
	//r := reflect.ValueOf(v)
	if v == nil {
		return nil, errors.New("param is required")
	}
	if !utils.IsPrimitive(kind) {

	}
	var result interface{}
	switch kind {
	case reflect.Bool:
		if r, ok := v.(bool); ok {
			result = r
		}
		break
	case reflect.Int:
		if r, ok := v.(int); ok {
			result = r
		}
		break
	case reflect.Int8:
		if r, ok := v.(int8); ok {
			result = r
		}
		break
	case reflect.Int16:
		if r, ok := v.(int16); ok {
			result = r
		}
		break
	case reflect.Int32:
		if r, ok := v.(int32); ok {
			result = r
		}
		break
	case reflect.Int64:
		if r, ok := v.(int64); ok {
			result = r
		}
		break
	case reflect.Float32:
		if r, ok := v.(float32); ok {
			result = r
		}
		break
	case reflect.Float64:
		if r, ok := v.(float64); ok {
			result = r
		}
		break
	case reflect.Uint:
		if r, ok := v.(uint); ok {
			result = r
		}
		break
	case reflect.Uint8:
		if r, ok := v.(uint8); ok {
			result = r
		}
		break
	case reflect.Uint16:
		if r, ok := v.(uint16); ok {
			result = r
		}
		break
	case reflect.Uint32:
		if r, ok := v.(uint32); ok {
			result = r
		}
		break
	case reflect.Uint64:
		if r, ok := v.(uint64); ok {
			result = r
		}
		break
	}

	if result == nil {
		if r, ok := v.(string); ok {
			switch kind {
			case reflect.Bool:
				result, err = strconv.ParseBool(r)
				break
			case reflect.Int:
				i, err := strconv.ParseInt(r, 10, 32)
				if err == nil {
					result = int(i)
				}
				break
			case reflect.Int8:
				i, err := strconv.ParseInt(r, 10, 8)
				if err == nil {
					result = int8(i)
				}
				break
			case reflect.Int16:
				i, err := strconv.ParseInt(r, 10, 16)
				if err == nil {
					result = int16(i)
				}
				break
			case reflect.Int32:
				i, err := strconv.ParseInt(r, 10, 32)
				if err == nil {
					result = int32(i)
				}
				break
			case reflect.Int64:
				i, err := strconv.ParseInt(r, 10, 64)
				if err == nil {
					result = int64(i)
				}
				break
			case reflect.Float32:
				i, err := strconv.ParseFloat(r, 32)
				if err == nil {
					result = float32(i)
				}
				break
			case reflect.Float64:
				i, err := strconv.ParseFloat(r, 64)
				if err == nil {
					result = float32(i)
				}
				break
			case reflect.Uint:
				i, err := strconv.ParseUint(r, 10, 32)
				if err == nil {
					result = uint(i)
				}
				break
			case reflect.Uint8:
				i, err := strconv.ParseUint(r, 10, 8)
				if err == nil {
					result = uint8(i)
				}
				break
			case reflect.Uint16:
				i, err := strconv.ParseUint(r, 10, 16)
				if err == nil {
					result = int16(i)
				}
				break
			case reflect.Uint32:
				i, err := strconv.ParseUint(r, 10, 32)
				if err == nil {
					result = uint32(i)
				}
				break
			case reflect.Uint64:
				i, err := strconv.ParseUint(r, 10, 64)
				if err == nil {
					result = uint64(i)
				}
				break
			}
		}
	}

	if err != nil {
		return nil, err
	}
	vall := reflect.ValueOf(result)
	return &vall, nil

}
