package opc

import "github.com/junhwong/mano/otpl/common"

type opBlock struct {
	opBase
}

func (op *opBlock) Load() (err error) {
	id, err := op.loader.ReadString()
	if err != nil {
		return
	}
	op.loader.PutBlock(id, op)
	return
}

func (op *opBlock) Exec(ctx common.Context) (common.Ptr, error) {
	ctx.Exec(op.loader, op.ptr+1)
	return common.ZeroPtr, nil
}

func init() {
	common.RegisterOpcode(0xC, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opBlock{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
