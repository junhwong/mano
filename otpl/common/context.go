package common

type TemplateFunc func(...interface{}) (interface{}, error)

// Context 表示一个运行时上下文
type Context interface {
	// 是否是严格模式
	IsStrict() bool
	// 从栈顶弹出一个元素，未找到返回 nil
	Pop() interface{}
	// 压入一个元素到栈顶
	Push(val interface{}) interface{}
	// 获取或设置一个变量
	Var(name string, val ...interface{}) (interface{}, bool)
	// 打印一个对象到输出
	Print(obj interface{}, escape bool) error
	// 启用一个新的域
	Scope() error
	// 注销一个域
	Unscope() error
	// 释放当前上下文
	Free()

	Load(src, ref string) (Loader, error)

	Exec(loader Loader, ptr Ptr) error

	TemplateFunc(name string, fn ...TemplateFunc) TemplateFunc
}
