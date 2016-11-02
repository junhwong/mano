package logs

import (
	"io"
	"runtime"
	"sync"
	"time"
)

type Entry struct {
	Time       time.Time
	Level      Level
	Message    []interface{}
	Caller     string
	Line       int
	Stack      []byte
	StackTrace []*runtime.Frame
}

type Provider interface {
	Output() io.Writer
	Log(entry *Entry)
}

type Logger struct {
	mu        sync.Mutex
	name      string
	level     Level
	provider  Provider
	logStack  bool
	calldepth int
}

func NewLogger(name string, provider ...Provider) *Logger {
	var p Provider
	if len(provider) > 0 && provider[0] != nil {
		p = provider[0]
	} else {
		p = GetProvider()
	}

	return &Logger{
		name:      name,
		level:     LERROR,
		provider:  p,
		calldepth: 2,
	}
}

func (l *Logger) Name() string {
	return l.name
}

func (l *Logger) Provider() Provider {
	return l.provider
}

func (l *Logger) Level() Level {
	return l.level
}

func (l *Logger) IsLogStack() bool {
	return l.logStack
}

func (l *Logger) isEnabled(level Level) bool {
	return l.level <= level
}

func (l *Logger) IsDebugEnabled() bool {
	return l.isEnabled(LDEBUG)
}

func (l *Logger) IsInfoEnabled() bool {
	return l.isEnabled(LINFO)
}

func (l *Logger) IsWarnEnabled() bool {
	return l.isEnabled(LWARN)
}

func (l *Logger) IsErrorEnabled() bool {
	return l.isEnabled(LERROR)
}

func (l *Logger) IsFatalEnabled() bool {
	return l.isEnabled(LFATAL)
}

func (l *Logger) log(level Level, message ...interface{}) {
	if !l.isEnabled(level) {
		return
	}

	entry := &Entry{
		Level:   level,
		Time:    time.Now(),
		Message: message,
	}

	if pc, _, line, ok := runtime.Caller(l.calldepth); ok {
		entry.Caller = runtime.FuncForPC(pc).Name()
		entry.Line = line
	} else {
		entry.Caller = "???"
		entry.Line = 0
	}

	if l.logStack || level >= LERROR {
		entry.StackTrace = getFrames(l.calldepth)
	}

	l.provider.Log(entry)
}

func (l *Logger) Debug(message ...interface{}) {
	l.log(LDEBUG, message...)
}

func (l *Logger) Info(message ...interface{}) {
	l.log(LINFO, message...)
}

func (l *Logger) Warn(message ...interface{}) {
	l.log(LWARN, message...)
}

func (l *Logger) Error(message ...interface{}) {
	l.log(LERROR, message...)
}

func (l *Logger) Fatal(message ...interface{}) {
	l.log(LFATAL, message...)
}

func (l *Logger) SetProvider(provider Provider) {
	if provider == nil {
		panic("[mano.logs.SetProvider] Invalid argument: provider")
	}

	l.mu.Lock()

	defer l.mu.Unlock()

	l.provider = provider

}

func (l *Logger) SetLevel(level Level) {

	l.mu.Lock()

	defer l.mu.Unlock()

	l.level = level

}

func (l *Logger) SetLogStack(logStack bool) {

	l.mu.Lock()

	defer l.mu.Unlock()

	l.logStack = logStack

}
