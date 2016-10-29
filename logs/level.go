package logs

type Level uint8

const (
	LDEBUG Level = iota
	LTRACE
	LINFO
	LWARN
	LERROR
	LFATAL
)

var labels = []string{"DEBUG", "TRACE", "INFO", "WARN", "ERROR", "FATAL"}

func (level Level) Lable() string {
	return labels[level]
}
