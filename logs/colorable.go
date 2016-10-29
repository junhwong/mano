package logs

import (
	"fmt"
	"strconv"
	"strings"
)

// ColorAttribute defines a single SGR Code
type ColorAttribute int

// Base attributes
const (
	Reset ColorAttribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack ColorAttribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack ColorAttribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack ColorAttribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack ColorAttribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

func ColorFormat(attrs ...ColorAttribute) string {
	format := make([]string, len(attrs))
	for i, v := range attrs {
		format[i] = strconv.Itoa(int(v))
	}
	return fmt.Sprintf("%s[%sm", "\x1b", strings.Join(format, ";"))
}

type ColorSetter func(isTerminal bool) string

var colorSetterMap = make(map[string]ColorSetter)

func GetColorSetter(name string) ColorSetter {
	setter, _ := colorSetterMap[name]
	return setter
}

func SetColorSetter(name string, setter ColorSetter) ColorSetter {
	colorSetterMap[name] = setter
	return GetColorSetter(name)
}

var ResetSetter = func() ColorSetter {
	color := ColorFormat(Reset)
	return func(isTerminal bool) string {
		if !isTerminal {
			return ""
		}
		return color
	}
}()

func init() {

	SetColorSetter("reset", ResetSetter)

	SetColorSetter("red", func() ColorSetter {
		color := ColorFormat(FgHiRed)
		return func(isTerminal bool) string {
			if !isTerminal {
				return ""
			}
			return color
		}
	}())

	SetColorSetter("cyan", func() ColorSetter {
		color := ColorFormat(FgCyan)
		return func(isTerminal bool) string {
			if !isTerminal {
				return ""
			}
			return color
		}
	}())

	SetColorSetter("green", func() ColorSetter {
		color := ColorFormat(FgHiGreen)
		return func(isTerminal bool) string {
			if !isTerminal {
				return ""
			}
			return color
		}
	}())

	SetColorSetter("blue", func() ColorSetter {
		color := ColorFormat(FgHiBlue)
		return func(isTerminal bool) string {
			if !isTerminal {
				return ""
			}
			return color
		}
	}())

}
