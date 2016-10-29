package opc

import (
	"github.com/junhwong/mano/otpl/common"
)

type opCastToIterator struct {
	opBase
}

func (op *opCastToIterator) Load() (err error) {
	return nil
}

func (op *opCastToIterator) Exec(ctx common.Context) (ptr common.Ptr, err error) {
	err = nil
	ptr = op.ptr + 1
	ctx.Push(newIterator(ctx.Pop(), ctx))
	return
}

func init() {
	common.RegisterOpcode(0x0F, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opCastToIterator{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
