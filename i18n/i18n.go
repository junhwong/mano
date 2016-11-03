package i18n

import (
	"io/ioutil"

	"github.com/junhwong/mano/logs"
)

type Localization interface {
	Local() string
	Lang(local string, name string, args ...interface{}) string
}

// type Bundles map[string]*Bundle

// func (b Bundles) Lang(local string, name string, args ...interface{}) string {
// 	local = strings.ToUpper(local)

// 	bundle, _ := b[local]
// 	if bundle != nil {
// 		index := strings.Index(local, "-")
// 		if index < 0 {
// 			index = strings.Index(local, "_")
// 		}
// 		if index > 0 {
// 			bundle, _ = b[local[:index]]
// 			if bundle == nil {
// 				bundle, _ = b[local[index+1:]]
// 			}
// 		}
// 	}

// 	if bundle == nil {
// 		return ""
// 	}

// 	return bundle.Lang(local, name, args...)
// }

func Load(dir string) (err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			panic(file.Name())
			logs.Debug(file.Name())
		}
	}
	return
}
