package opc

import "github.com/junhwong/mano/otpl/common"

type opNop struct {
	opBase
}

func (op *opNop) Load() (err error) {
	return nil
}

func (op *opNop) Exec(ctx common.Context) (common.Ptr, error) {
	return op.ptr + 1, nil
}

func init() {
	common.RegisterOpcode(0x02, func(loader common.Loader, ptr common.Ptr, line common.LineNo, flag byte) common.Opcode {

		return &opNop{
			opBase: opBase{
				loader:     loader,
				ptr:        ptr,
				lineNumber: line,
				flag:       flag,
			},
		}
	})
}
