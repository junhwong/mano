package runtime

import (
	"crypto/md5"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/junhwong/mano/otpl/common"
)

func canonicalName(template, base string) string {
	template = strings.Trim(template, " ")
	base = strings.Replace(strings.Trim(base, " "), "\\", "/", -1)

	if last := strings.LastIndexAny(base, "/"); last > -1 {
		base = base[:last]
	}

	if !(len(base) > 0 && base[0] == '/') {
		base = "/" + base
	}

	if !(len(template) > 0 && template[0] == '/') {
		base = "/"
	}

	template = strings.Replace(filepath.Join(base, template), "\\", "/", -1)

	return template
}

func md5Sign(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func NewRuntime(ilpath string, debug bool) *Runtime {
	return &Runtime{
		ilpath: ilpath,
		debug:  debug,
	}
}

type Runtime struct {
	ilpath        string
	debug         bool
	strictMode    bool
	allowedSuffix []string
}

func (rt *Runtime) ILPath() string {
	return rt.ilpath
}

func (rt *Runtime) Context(data map[string]interface{}, out io.Writer) common.Context {

	ctx := &context{
		currentScope: newContextScope(data, nil),
		out:          out,
		runtime:      rt,
		loaders:      make(map[string]common.Loader, 0),
		scopes:       make([]contextScope, 0), //make(common.Stack, 0),
	}

	return ctx
}

func (rt *Runtime) exec() string {
	return rt.ilpath
}

func (rt *Runtime) Render(data map[string]interface{}, path string, out io.Writer) (err error) {
	ctx := rt.Context(data, out)
	defer ctx.Free()

	loader, err := ctx.Load(path, "")

	if err != nil {
		return
	}

	err = ctx.Exec(loader, loader.StartPtr())
	if err != nil {
		return
	}
	return
}
