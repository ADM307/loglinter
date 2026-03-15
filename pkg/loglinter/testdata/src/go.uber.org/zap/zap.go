package zap

import "go.uber.org/zap/zapcore"

type Logger struct{}

func (l *Logger) Info(msg string, fields ...Field)  {}
func (l *Logger) Error(msg string, fields ...Field) {}
func (l *Logger) Debug(msg string, fields ...Field) {}
func (l *Logger) Warn(msg string, fields ...Field)  {}

type Field = zapcore.Field

func String(key string, val string) Field {
	return zapcore.Field{Key: key, String: val, Type: zapcore.StringType}
}

func L() *Logger {
	return &Logger{}
}
