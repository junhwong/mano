package opc

import (
	"fmt"

	"github.com/junhwong/mano/logs"
	"github.com/junhwong/mano/otpl/common"
)

const (
	INCLUDE = 0x01
	REQUIRE = 0x02
	LAYOUT  = 0x03
)

type opReference struct {
	opBase
	src string
}

func (op *opReference) Load() (err error) {
	op.src, err = op.loader.ReadString()

	return
}

func (op *opReference) Exec(ctx common.Context) (ptr common.Ptr, err error) {
	err = nil
	ptr = op.ptr + 1

	loader, err := ctx.Load(op.src, op.loader.TemplateName())
	if err != nil {
		return
	}
	switch op.flag {
	case INCLUDE:
		if loader != nil {
			err = ctx.Exec(loader, loader.StartPtr())
		} else {
			logs.Debug("Failed to loding (or not found) include template: %s", op.src)
		}
		break
	case REQUIRE:
		if loader == nil {
			err = fmt.Errorf("Failed to loding (or not found) require template: %s", op.src)
		}
		//nothing to do,its only import headers
		break
	case LAYOUT:

		if loader == nil {
			err = fmt.Errorf("Failed to loding (or not found) layout template: %s", op.src)
		} else {
			loader.SetBody(op.loader, ptr)
			err = ctx.Exec(loader, loader.StartPtr())
		}
		ptr = common.ZeroPtr
		return
	default:
		err = fmt.Errorf("Undefined reference type: 0x%X", op.flag)
		break
	}

	return
}

func init() {
	common.RegisterOpcode(0x0E, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opReference{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
