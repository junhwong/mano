package opc

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/junhwong/mano/otpl/common"
)

type Operator byte

const (
	ADD Operator = 0x01
	SUB Operator = 0x02
	MUL Operator = 0x03
	DIV Operator = 0x04
	MOD Operator = 0x05
	//取负
	NEG Operator = 0x06
	//取正
	POS Operator = 0x07
	EQ  Operator = 0x08
	NE  Operator = 0x09
	GT  Operator = 0x0A
	//逻辑大于等于
	GE Operator = 0x0B
	LT Operator = 0x0C
	//逻辑小于等于
	LE  Operator = 0x0D
	AND Operator = 0x0E
	OR  Operator = 0x0F
)

type opOperation struct {
	opBase
	// operator Operator
}

func (op *opOperation) Load() (err error) {
	return nil
}

func hasFloat(v ...common.Number) bool {
	for _, num := range v {
		if num.IsFloat() {
			return true
		}
	}
	return false
}

func getNum(a, b interface{}) (common.Number, common.Number, error) {
	ar, err := common.ToNumber(a)
	if err != nil {
		return nil, nil, err
	}
	br, err := common.ToNumber(b)
	if err != nil {
		return nil, nil, err
	}
	return ar, br, nil
}

func getBool(v interface{}) (bool, error) {
	if b, ok := v.(bool); ok {
		return b, nil
	} else if s, ok := v.(string); ok {
		return strconv.ParseBool(s)
	}
	return false, fmt.Errorf("转换为bool失败：%v", reflect.TypeOf(v))
}

func (op *opOperation) Exec(ctx common.Context) (ptr common.Ptr, err error) {
	var result interface{}
	ptr = 0
	err = nil
	operator := Operator(op.flag)
	switch operator {
	case ADD:
		b, a, err := getNum(ctx.Pop(), ctx.Pop())
		if err != nil {
			break
		}
		r, err := common.ToNumber(a.Float() + b.Float())
		if err != nil {
			break
		}
		if hasFloat(a, b) {
			result = r.Float()
		} else {
			result = r.Int()
		}

		break
	case SUB:
		b, a, err := getNum(ctx.Pop(), ctx.Pop())
		if err != nil {
			break
		}
		r, err := common.ToNumber(a.Float() - b.Float())
		if err != nil {
			break
		}
		if hasFloat(a, b) {
			result = r.Float()
		} else {
			result = r.Int()
		}
		break
	case MUL:
		b, a, err := getNum(ctx.Pop(), ctx.Pop())
		if err != nil {
			break
		}
		r, err := common.ToNumber(a.Float() * b.Float())
		if err != nil {
			break
		}
		if hasFloat(a, b) {
			result = r.Float()
		} else {
			result = r.Int()
		}
		break
	case DIV:
		b, a, err := getNum(ctx.Pop(), ctx.Pop())
		if err != nil {
			break
		}
		r, err := common.ToNumber(a.Float() / b.Float())
		if err != nil {
			break
		}
		if hasFloat(a, b) {
			result = r.Float()
		} else {
			result = r.Int()
		}
		break
	case MOD:
		b, a, err := getNum(ctx.Pop(), ctx.Pop())
		if err != nil {
			break
		}
		r, err := common.ToNumber(a.Int() % b.Int()) //float64 不能用于 % 操作
		if err != nil {
			break
		}
		if hasFloat(a, b) {
			result = r.Float()
		} else {
			result = r.Int()
		}
		break
	case NEG:
		a, err := common.ToNumber(ctx.Pop())
		if err != nil {
			break
		}

		if hasFloat(a) {
			result = +(a.Float())
		} else {
			result = +(a.Int())
		}
		break
	case POS:
		a, err := common.ToNumber(ctx.Pop())
		if err != nil {
			break
		}
		if hasFloat(a) {
			result = -(a.Float())
		} else {
			result = -(a.Int())
		}
		break
	case EQ:
		b, a, err := getNum(ctx.Pop(), ctx.Pop())
		if err != nil {
			break
		}
		result = a.Float() == b.Float()
		break
	case NE:
		b, a, err := getNum(ctx.Pop(), ctx.Pop())
		if err != nil {
			break
		}
		result = a.Float() != b.Float()
		break
	case GT:
		b, a, err := getNum(ctx.Pop(), ctx.Pop())
		if err != nil {
			break
		}
		result = a.Float() > b.Float()
		break
	case GE:
		b, a, err := getNum(ctx.Pop(), ctx.Pop())
		if err != nil {
			break
		}
		result = a.Float() >= b.Float()
		break
	case LT:
		b, a, err := getNum(ctx.Pop(), ctx.Pop())
		if err != nil {
			break
		}

		result = a.Float() < b.Float()
		break
	case LE:
		b, a, err := getNum(ctx.Pop(), ctx.Pop())
		if err != nil {
			break
		}
		result = a.Float() <= b.Float()
		break
	case AND:
		b, err := getBool(ctx.Pop())
		if err != nil {
			break
		}
		a, err := getBool(ctx.Pop())
		if err != nil {
			break
		}
		result = a && b
		break
	case OR:
		b, err := getBool(ctx.Pop())
		if err != nil {
			break
		}
		a, err := getBool(ctx.Pop())
		if err != nil {
			break
		}
		result = a && b
		break
	default:
		err = fmt.Errorf("undefined operator:0x%X", operator)
		break
	}

	ptr = op.ptr + 1
	if err != nil && !ctx.IsStrict() {
		err = nil
		result = nil
	}

	ctx.Push(result)
	return
}

func init() {
	common.RegisterOpcode(0x09, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opOperation{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
