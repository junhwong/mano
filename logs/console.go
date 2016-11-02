package logs

import (
	"fmt"
	"io"

	"github.com/mattn/go-colorable"
)

type Console struct {
	output io.Writer
}

func NewConsole() *Console {
	return &Console{
		output: colorable.NewColorableStdout(),
	}
}

func (c *Console) Output() io.Writer {
	return c.output
}

func (c *Console) Log(entry *Entry) {

	// entry := &Entry{
	// 	Level:      level,
	// 	IsTerminal: true,
	// 	Time:       time.Now(),
	// 	Message:    message,
	// 	Stack:      debug.Stack(),
	// }

	msg := FormatLog(true, entry)
	_, err := fmt.Fprint(c.output, msg)
	if err != nil {
		panic(fmt.Errorf("[mano.logs.Console.Log] %v", err))
	}

}
