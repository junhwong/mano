package opc

import "github.com/junhwong/mano/otpl/common"

type opLoadVariable struct {
	opBase
	name string
}

func (op *opLoadVariable) Load() (err error) {
	op.name, err = op.loader.ReadString()

	return
}

func (op *opLoadVariable) Exec(ctx common.Context) (ptr common.Ptr, err error) {
	ptr = op.ptr + 1
	val, ok := ctx.Var(op.name)
	err = handErr(!ok, ctx, "unset variable:%s", op.name)
	if err == nil {
		ctx.Push(val)
	}
	return
}

func init() {
	common.RegisterOpcode(0x05, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opLoadVariable{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
