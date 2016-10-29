package otpl

import "github.com/junhwong/mano/otpl/runtime"

func New(ilpath string) *runtime.Runtime {
	return runtime.NewRuntime(ilpath, true)
}

// func (inter *Interpreter) Render(data map[string]interface{}, path string, out io.Writer) {

// }

// func Default(data map[string]interface{}) {

// }
