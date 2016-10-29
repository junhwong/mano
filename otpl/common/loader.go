package common

type Loader interface {
	ReadByte() (byte, error)
	ReadBool() (bool, error)
	ReadPtr() (Ptr, error)
	ReadInt() (int, error)
	ReadLong() (int64, error)
	ReadFloat() (float64, error)
	ReadString() (string, error)
	Load(ptr Ptr) (Opcode, error)

	TemplateName() string
	StartPtr() Ptr

	PutBlock(blockID string, block Opcode)
	GetBlock(blockID string) (block Opcode, ok bool)

	SetBody(body Loader, start Ptr)
	BodyLoader() (Loader, bool)
	BodyPtr() Ptr

	Close() error

	Blocks() map[string]Opcode
}
