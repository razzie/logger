package logger

import (
	"io"
	"log"
)

// Logger is an interface which is compatible with log.Logger
// (but without the useless and harmful Fatal and Panic functions)
type Logger interface {
	Flags() int
	SetFlags(flag int)
	Prefix() string
	SetPrefix(prefix string)
	Writer() io.Writer
	SetOutput(w io.Writer)
	Output(calldepth int, s string) error
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// NilLogger returns a Logger which uses the default log.Logger internally
func NilLogger() Logger {
	return &nilLogger{}
}

type nilLogger struct{}

func (nilLogger) Flags() int {
	return log.Flags()
}

func (nilLogger) SetFlags(flag int) {
	log.SetFlags(flag)
}

func (nilLogger) Prefix() string {
	return log.Prefix()
}

func (nilLogger) SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}

func (nilLogger) Writer() io.Writer {
	return log.Writer()
}

func (nilLogger) SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func (nilLogger) Output(calldepth int, s string) error {
	return log.Output(calldepth+1, s)
}

func (nilLogger) Print(v ...interface{}) {
	log.Print(v...)
}

func (nilLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (nilLogger) Println(v ...interface{}) {
	log.Println(v...)
}
