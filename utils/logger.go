package utils

import "go.uber.org/zap"

type Log interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type Logger struct {
	ZapLogger *zap.SugaredLogger
}

func NewZapLogger(zapLogger *zap.SugaredLogger) *Logger {
	return &Logger{ZapLogger: zapLogger}
}

func (log *Logger) Info(args ...interface{}) {
	log.ZapLogger.Info(args)
}

func (log *Logger) Error(args ...interface{}) {
	log.ZapLogger.Error(args)
}
