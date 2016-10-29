package opc

import "github.com/junhwong/mano/otpl/common"

type opScope struct {
	opBase
	mode bool
}

func (op *opScope) Load() (err error) {
	if op.flag == 0x00 {
		op.mode = false
	} else {
		op.mode = true
	}

	return
}

func (op *opScope) Exec(ctx common.Context) (ptr common.Ptr, err error) {
	err = nil
	ptr = op.ptr + 1

	if op.mode {
		ctx.Scope()
	} else {
		ctx.Unscope()
	}
	return
}

func init() {
	common.RegisterOpcode(0x0B, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opScope{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
