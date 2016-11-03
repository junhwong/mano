package i18n

type MinorBundle struct {
	local string
	langs map[string]string
}

func (mb *MinorBundle) Local() string {
	return mb.local
}

type Bundle struct {
	MinorBundle
	Minors map[string]*MinorBundle
}

func (b *Bundle) Lang(minorLocal, name string) (lang string) {
	var ok bool

	// 获取指定语言的资源
	minorBundle, _ := b.Minors[minorLocal]
	if minorBundle != nil {
		if lang, ok = minorBundle.langs[name]; ok {
			return
		}
	}
	// 获取主语言的资源
	if lang, ok = b.langs[name]; ok {
		return
	}
	// 获取其它存在的资源
	for _, minorBundle = range b.Minors {
		if lang, ok = minorBundle.langs[name]; ok {
			return
		}
	}
	return
}
