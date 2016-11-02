package logs

import "runtime"

type Exception struct {
	message []interface{}
	stack   []byte
	cause   error
	frames  []*runtime.Frame
}

func (ex *Exception) Message() []interface{} {
	return ex.message
}

func (ex *Exception) Stack() []byte {
	return ex.stack
}

func (ex *Exception) Trace() []*runtime.Frame {
	return ex.frames
}

func (ex *Exception) Cause() error {
	return ex.cause
}

func (ex *Exception) Error() string {
	if len(ex.message) == 0 {
		return ex.cause.Error()
	}
	panic(formatMessage(false, ex.message))
	return formatMessage(false, ex.message)
}

func getFrames(calldepth int) []*runtime.Frame {
	var frames []*runtime.Frame
	for {
		if pc, file, line, ok := runtime.Caller(calldepth); ok {
			frames = append(frames, &runtime.Frame{
				File:     file,
				Line:     line,
				Function: runtime.FuncForPC(pc).Name(),
			})
		} else {
			break
		}
		calldepth++
	}
	return frames
}

// NewError returns a Exception with the specified detail message and cause.
// cause is a optional
func NewError(cause error, message ...interface{}) *Exception {

	if cause == nil && len(message) == 0 {
		panic("[mano.logs.NewError] Invalid argument")
	}
	frames := getFrames(2)
	return &Exception{
		message: message,
		cause:   cause,
		frames:  frames,
		//stack:   debug.Stack(),
	}
}
