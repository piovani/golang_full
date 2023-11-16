package logger

import (
	"fmt"
	"log/slog"
)

type Logger struct {
	prefix string
	log    *slog.Logger
}

func NewLogger(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
		log:    slog.New(slog.Default().Handler()),
	}
}

func (l *Logger) Info(msg string, event ...any) {
	if l.prefix == "" {
		l.log.Info(msg, event...)
	} else {
		l.log.Info(fmt.Sprintf("%s: %s", l.prefix, msg), event...)
	}
}

func (l *Logger) Error(err error, event ...any) {
	if l.prefix == "" {
		l.log.Info(err.Error(), event...)
	} else {
		l.log.Info(fmt.Sprintf("%s: %s", l.prefix, err.Error()), event...)
	}
}
