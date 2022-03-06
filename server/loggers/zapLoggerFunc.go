package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type MyLogger interface {
	Fatalf(string, ...interface{})
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Debug(...interface{})
	Warn(...interface{})
	Info(...interface{})
	Fatal(...interface{})
}

// PrintfLogger for fasthttp log interface

type PrintfLogger interface {
	Printf(format string, args ...interface{})
}

// Debug logs a messageat level DebugMode on the ZapLogger
func (l *ZapLogger) Debug(args ...interface{}) {
	l.Log.Debug(fmt.Sprint(args...))
}

// Debugf logs a message at level DebugMode on the ZapLogger
func (l *ZapLogger) Debugf(template string, args ...interface{}) {
	l.Log.Debug(fmt.Sprintf(template, args...))
}

// Info logs a message at level Info on the ZapLogger
func (l *ZapLogger) Info(args ...interface{}) {
	l.Log.Info(fmt.Sprint(args...))
}

// Infof logs a message at level Info on the ZapLogger
func (l *ZapLogger) Infof(template string, args ...interface{}) {
	l.Log.Info(fmt.Sprintf(template, args...))
}

// Warn logs a message at level Warn on the ZapLogger
func (l *ZapLogger) Warn(args ...interface{}) {
	l.Log.Warn(fmt.Sprint(args...))
}

// Warning logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Warning(args ...interface{}) {
	l.Log.Warn(fmt.Sprint(args...))
}

// Warnf logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Warnf(template string, args ...interface{}) {
	l.Log.Warn(fmt.Sprintf(template, args...))
}

// Warningf logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Warningf(template string, args ...interface{}) {
	l.Log.Warn(fmt.Sprintf(template, args...))
}

// Error logs a message at level Error on the ZapLogger.
func (l *ZapLogger) Error(args ...interface{}) {
	l.Log.Error(fmt.Sprint(args...))
}

// Errorf logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Errorf(template string, args ...interface{}) {
	l.Log.Error(fmt.Sprintf(template, args...))
}

// Fatal logs a message at level Fatal on the ZapLogger.
func (l *ZapLogger) Fatal(args ...interface{}) {
	l.Log.Fatal(fmt.Sprint(args...))
}

// Fatalf logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Fatalf(template string, args ...interface{}) {
	l.Log.Fatal(fmt.Sprintf(template, args...))
}

// Panic logs a message at level Painc on the ZapLogger.
func (l *ZapLogger) Panic(args ...interface{}) {
	l.Log.Panic(fmt.Sprint(args...))
}

// DPanic logs a message at level Painc on the ZapLogger.
func (l *ZapLogger) DPanic(args ...interface{}) {
	l.Log.DPanic(fmt.Sprint(args...))
}

// Panicf  logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Panicf(template string, args ...interface{}) {
	l.Log.Panic(fmt.Sprintf(template, args...))
}

// Print logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Print(args ...interface{}) {
	l.Log.Info(fmt.Sprint(args...))
}

// With return a log with an extra field.
func (l *ZapLogger) With(k string, v interface{}) *ZapLogger {
	l.Log.With(zap.Any(k, v))
	return l
}

// WithField return a log with an extra field.
func (l *ZapLogger) WithField(k string, v interface{}) *ZapLogger {
	l.Log.With(zap.Any(k, v))
	return l
}

// WithFields return a log with extra fields.
func (l *ZapLogger) WithFields(fields map[string]interface{}) *ZapLogger {
	clog := l
	i := 0
	for k, v := range fields {
		if i == 0 {
			clog = l.WithField(k, v)
		} else {
			clog = clog.WithField(k, v)
		}
		i++
	}
	return clog
}
