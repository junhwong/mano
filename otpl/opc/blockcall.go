package opc

import (
	"errors"
	"strings"

	"github.com/junhwong/mano/otpl/common"
)

type opBlockCall struct {
	opBase
	id string
	// parameterCount int
}

func (op *opBlockCall) Load() (err error) {

	op.id, err = op.loader.ReadString()

	return
}

func (op *opBlockCall) Exec(ctx common.Context) (ptr common.Ptr, err error) {
	err = nil
	ptr = op.ptr + 1
	// parameterCount:=op.flag
	loader := op.loader

	if strings.EqualFold("body", op.id) {
		if bodyLoader, ok := loader.BodyLoader(); ok {
			err = ctx.Exec(bodyLoader, loader.BodyPtr())
		} else {
			err = errors.New("Invalid operation,tags: @body") // handErr(ok, ctx, "布局页面不能直接执行，标签：@body")
		}
	} else if block, ok := loader.GetBlock(op.id); ok {
		block.Exec(ctx)
	}
	// ignored other case,if block is ni

	return
}

func init() {
	common.RegisterOpcode(0x0D, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opBlockCall{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
