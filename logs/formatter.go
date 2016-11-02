package logs

import (
	"fmt"
	"regexp"
	"strings"
)

var reg = regexp.MustCompile(`\{\s?\w+\s?\:`)

func formatMessage(isTerminal bool, v ...interface{}) string {
	msg := ""

	if len(v) > 0 {
		if src, ok := v[0].(string); ok {
			format := ""
			for {
				loc := reg.FindStringIndex(src)
				if loc != nil && len(loc) == 2 {
					index := strings.Index(src[loc[1]:], "}")
					if index < 0 {
						format += src[0:loc[1]]
						src = src[loc[1]:]
					} else {
						temp := strings.Trim(src[loc[0]+1:loc[1]], " ")
						temp = strings.Trim(temp, ":")
						temp = strings.Trim(temp, " ")

						format += src[:loc[0]]
						if temp != "" {
							if setter := GetColorSetter(temp); setter != nil {
								format += setter(isTerminal)
							} else {
								//todo
							}
						}

						format += src[loc[1] : loc[1]+index]
						format += ResetSetter(isTerminal)
						src = src[loc[1]+index+1:]
					}
				} else {
					format += src
					src = ""
					break
				}
			}
			msg += fmt.Sprintf(format, v[1:]...)

		} else if err, ok := v[0].(error); ok {
			msg += err.Error()
		} else {
			msg += fmt.Sprint("%v", v[0])
		}
	}
	return msg
}

func getFilename(file string) string {
	if file == "" {
		return "???"
	}
	last := strings.LastIndex(file, "/")
	if last < 0 {
		last = strings.LastIndex(file, "\\")
	}
	if last < 0 {
		return file
	}
	return file[last+1:]
}

func FormatLog(isTerminal bool, entry *Entry) string {
	if entry == nil {
		return ""
	}

	msg := entry.Time.Format("2006-01-02 15:04:05")
	msg += " "
	// msg += GetColorSetter("cyan")(isTerminal)
	// msg += entry.Caller
	// msg += ResetSetter(isTerminal)
	msg += fmt.Sprintf("%s%s:%d%s\n", GetColorSetter("cyan")(isTerminal), entry.Caller, entry.Line, ResetSetter(isTerminal))
	// msg += "\n"
	msg += fmt.Sprintf("%s%s%s:", colorForLevel(entry.Level, isTerminal), entry.Level.Lable(), ResetSetter(isTerminal))

	if ex, ok := entry.Message[0].(*Exception); ok {
		msg += ex.Error()
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		i := 0
		max := len(ex.Trace())
		for i < max {
			frame := ex.Trace()[i]
			if strings.EqualFold(frame.Function, "runtime.goexit") {
				break
			}

			msg += fmt.Sprintf("\tat %s(%s:%d)\n", frame.Function, getFilename(frame.File), frame.Line)

			i++
		}

		//msg += string(ex.Stack())
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}

	} else {
		msg += formatMessage(isTerminal, entry.Message...)
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}

		if entry.StackTrace != nil && len(entry.StackTrace) > 0 {
			i := 0
			max := len(entry.StackTrace)
			for i < max {
				frame := entry.StackTrace[i]
				if strings.EqualFold(frame.Function, "runtime.goexit") {
					break
				}

				msg += fmt.Sprintf("\tat %s(%s:%d)\n", frame.Function, getFilename(frame.File), frame.Line)

				i++
			}
			if !strings.HasSuffix(msg, "\n") {
				msg += "\n"
			}
		}
	}

	return msg
}

// var (
// 	green   = "" // "\033[34;1m"
// 	white   = "" // string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
// 	yellow  = "" // "\033[33;1m"
// 	red     = "" // "\033[31;1m"
// 	blue    = "\033[32;1m"
// 	magenta = "" // string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
// 	cyan    = "" // string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
// 	reset   = "\033[0m"
// )

func colorForLevel(level Level, isTerminal bool) string {
	var setter ColorSetter
	switch level {
	case LINFO:
		setter = GetColorSetter("white")
	case LDEBUG:
		setter = GetColorSetter("blue")
	case LWARN:
		setter = GetColorSetter("yello")
	case LERROR:
		setter = GetColorSetter("red")
	}

	//TODO
	setter = GetColorSetter("green")

	if setter != nil {
		return setter(isTerminal)
	} else if isTerminal {
		return ResetSetter(isTerminal)
	}

	return ""
}
