package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogLevel zap log level
// var defaultLevel = zap.NewAtomicLevel() // zap.NewAtomicLevelAt(zap.DebugLevel)
var _defaultLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)

// SetLevel set log level for zap log
func SetLevel(level zapcore.Level) {
	_defaultLevel.SetLevel(level)
}

// SetLevel set log level for zap log
func (l *Logger) SetLevel(level zapcore.Level) {
	_defaultLevel.SetLevel(level)
}
