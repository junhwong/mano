package opc

import (
	"errors"
	"fmt"

	"github.com/junhwong/mano/otpl/common"
)

const (
	typeEXIT  = 0x01
	typeNEVER = 0x02
	typeTRUE  = 0x03
	typeFALSE = 0x04
)

type opJump struct {
	opBase
	target common.Ptr
}

func (op *opJump) Load() (err error) {
	switch op.flag {
	case typeEXIT:
		break
	case typeFALSE, typeNEVER, typeTRUE:
		op.target, err = op.loader.ReadPtr()
		break
	default:
		err = fmt.Errorf("undefined jump type：0x%X", op.flag)
		break
	}

	return
}

func toBool(v interface{}) bool {

	if b, ok := v.(bool); ok {
		return b
	} else if b, ok := v.(float64); ok {
		return b > 0
	}
	return v != nil
}

func (op *opJump) Exec(ctx common.Context) (common.Ptr, error) {

	switch op.flag {
	case typeEXIT:
		return common.ZeroPtr, nil
	case typeNEVER:
		return op.target, nil
	case typeFALSE:

		val := toBool(ctx.Pop())
		if !val {
			return op.target, nil
		}

		break
	case typeTRUE:
		val := toBool(ctx.Pop())
		if val {
			return op.target, nil
		}
		break
	default:
		return common.ZeroPtr, errors.New("Invalid jump flags")
	}

	return op.ptr + 1, nil
}

func init() {
	common.RegisterOpcode(0x03, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opJump{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
