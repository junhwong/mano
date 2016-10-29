package opc

import "github.com/junhwong/mano/otpl/common"

type opPrint struct {
	opBase
	escape bool
}

func (op *opPrint) Load() (err error) {

	if op.flag == 0x00 {
		op.escape = false
	} else {
		op.escape = true
	}

	return
}

func (op *opPrint) Exec(ctx common.Context) (common.Ptr, error) {

	err := ctx.Print(ctx.Pop(), op.escape)

	return op.ptr + 1, err
}

func init() {
	common.RegisterOpcode(0x08, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opPrint{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
