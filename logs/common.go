package logs

import (
	"os"
	"strings"
)

var rootLogger *Logger

func init() {
	rootLogger = NewLogger("root", NewConsole())
	rootLogger.calldepth = 3
}

/*****************************************/

func GetProvider() Provider {
	return rootLogger.Provider()
}

func GetLevel() Level {
	return rootLogger.Level()
}

func IsLogStack() bool {
	return rootLogger.IsLogStack()
}

func SetProvider(provider Provider) {
	rootLogger.SetProvider(provider)
}

func SetLevel(level Level) {
	rootLogger.SetLevel(level)
}

func SetLogStack(logStack bool) {
	rootLogger.SetLogStack(logStack)
}

/*****************************************/

func Debug(message ...interface{}) {
	rootLogger.Debug(message...)
}

func Info(message ...interface{}) {
	rootLogger.Info(message...)
}

func Warn(message ...interface{}) {
	rootLogger.Warn(message...)
}

func Error(message ...interface{}) {
	rootLogger.Error(message...)
}

func Fatal(message ...interface{}) {
	rootLogger.Fatal(message...)
}

/*****************************************/

func IsDebugEnabled() bool {
	return rootLogger.IsDebugEnabled()
}

func IsInfoEnabled() bool {
	return rootLogger.IsDebugEnabled()
}

func IsWarnEnabled() bool {
	return rootLogger.IsDebugEnabled()
}

func IsErrorEnabled() bool {
	return rootLogger.IsDebugEnabled()
}

func IsFatalEnabled() bool {
	return rootLogger.IsDebugEnabled()
}

/*****************************************/

func Print(v ...interface{}) {
	os.Stdout.WriteString(formatMessage(true, v...))
}

func Println(v ...interface{}) {
	msg := formatMessage(true, v...)
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}
	os.Stdout.WriteString(msg)
}
