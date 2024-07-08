package logger

import (
	"go.uber.org/zap"
)

type ILogger interface {
	Info(msg string, fields ...Field)
	Warning(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	With(fields ...Field) ILogger
}

type logger struct {
	zap *zap.Logger
}

func (l logger) Info(msg string, fields ...Field) {
	l.zap.Info(msg, fields...)
}

func (l logger) Error(msg string, fields ...Field) {
	l.zap.Error(msg, fields...)
}

func (l logger) Warning(msg string, fields ...Field) {
	l.zap.Warn(msg, fields...)
}
func (l logger) Fatal(msg string, fields ...Field) {
	l.zap.Fatal(msg, fields...)
}

func (l logger) With(fields ...Field) ILogger {
	clonedLogger := l.zap.With(fields...)
	return logger{zap: clonedLogger}
}

func New(namespace string) ILogger {
	return logger{
		zap: newZapLogger(namespace),
	}
}
