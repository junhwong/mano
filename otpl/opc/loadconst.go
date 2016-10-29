package opc

import (
	"fmt"

	"github.com/junhwong/mano/otpl/common"
)

const (
	NULL  = 0x00
	STR   = 0x01
	INT   = 0x02
	LONG  = 0x03
	FLOAT = 0x04
	TRUE  = 0x05
	FLASE = 0x06
)

type opLoadConst struct {
	opBase
	value interface{}
}

func (op *opLoadConst) Load() (err error) {
	switch op.flag {
	case NULL:
		op.value = nil
		break
	case STR:
		op.value, err = op.loader.ReadString()
		break
	case INT:
		op.value, err = op.loader.ReadInt()
		break
	case LONG:
		op.value, err = op.loader.ReadLong()
		break
	case FLOAT:
		op.value, err = op.loader.ReadFloat()
		break
	case TRUE:
		op.value = true
		break
	case FLASE:
		op.value = false
		break
	default:
		err = fmt.Errorf("Undefined const type：0x%X", op.flag)
	}

	return
}

func (op *opLoadConst) Exec(ctx common.Context) (common.Ptr, error) {
	ctx.Push(op.value)
	return op.ptr + 1, nil
}

func init() {
	common.RegisterOpcode(0x04, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opLoadConst{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
