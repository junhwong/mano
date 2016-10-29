package opc

import "github.com/junhwong/mano/otpl/common"

type opSetVariable struct {
	opBase
	name string
}

func (op *opSetVariable) Load() (err error) {
	op.name, err = op.loader.ReadString()

	return
}

func (op *opSetVariable) Exec(ctx common.Context) (common.Ptr, error) {
	ctx.Var(op.name, ctx.Pop())
	return op.ptr + 1, nil
}

func init() {
	common.RegisterOpcode(0x06, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opSetVariable{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
