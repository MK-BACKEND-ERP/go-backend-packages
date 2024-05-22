package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
)

func NewLogger(prefix ...string) *Logger {
	pre, preFormated := "", ""

	writer := io.Writer(os.Stdout)

	if len(prefix) > 0 && prefix[0] != "" {
		pre = prefix[0]
		preFormated = "(" + prefix[0] + "): "
	}

	logger := log.New(writer, pre, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG "+preFormated, logger.Flags()),
		info:    log.New(writer, blue+"INFO "+reset+preFormated, logger.Flags()),
		warning: log.New(writer, yellow+"WARNING "+reset+preFormated, logger.Flags()),
		err:     log.New(writer, red+"ERROR "+reset+preFormated, logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Debug(v ...any) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...any) {
	l.info.Println(v...)
}
func (l *Logger) Warn(v ...any) {
	l.warning.Println(v...)
}
func (l *Logger) Error(v ...any) {
	l.err.Println(v...)
}

func (l *Logger) Debugf(format string, v ...any) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}
func (l *Logger) Warnf(format string, v ...any) {
	l.warning.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...any) {
	l.err.Printf(format, v...)
}

func (l *Logger) ErrorFatal(v ...any) {
	l.err.Fatal(v...)
}
