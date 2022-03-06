package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger zap log
type ZapLogger struct {
	path       string
	prefix     string
	Log        *zap.Logger
	debug      bool
	storeInDay bool
	addCaller  bool

	logLevel zapcore.LevelEnabler
	days     int64
}

// Logger  logger
type Logger = ZapLogger

// An Option configures a Logger.
type Option interface {
	apply(*Logger)
}

// optionFunc wraps a func so it satisfies the Option interface.
type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) {
	f(log)
}

func (l *Logger) clone() *Logger {
	copy := *l
	return &copy
}

// WithOptions clones the current Logger, applies the supplied Options, and
// returns the resulting Logger. It's safe to use concurrently.
func (l *Logger) WithOptions(opts ...Option) *Logger {
	c := l.clone()
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}

// AddCaller add caller
func AddCaller() Option {
	return optionFunc(func(log *Logger) {
		log.addCaller = true
	})
}

// ZapLoggerOption options
type ZapLoggerOption func(*ZapLogger)

// WithDebug set debug
func WithDebug() ZapLoggerOption {
	return func(zl *ZapLogger) {
		zl.debug = true
	}
}

// WithDays set debug
func WithDays(days int64) ZapLoggerOption {
	return func(zl *ZapLogger) {
		zl.storeInDay = true
		zl.days = days
	}
}

// WithStoreInDay set debug
func WithStoreInDay() ZapLoggerOption {
	return func(zl *ZapLogger) {
		zl.storeInDay = true
	}
}

// WithCaller set debug
func WithCaller() ZapLoggerOption {
	return func(zl *ZapLogger) {
		zl.addCaller = true
	}
}

// WithLevel set log level
func WithLevel(level zapcore.Level) ZapLoggerOption {
	return func(zl *ZapLogger) {
		_defaultLevel.SetLevel(level)
		zl.logLevel = _defaultLevel
	}
}

// WithPath set log level
func WithPath(path string) ZapLoggerOption {
	return func(zl *ZapLogger) {
		zl.path = path
	}
}

// WithPrefix set log level
func WithPrefix(prefix string) ZapLoggerOption {
	return func(zl *ZapLogger) {
		zl.prefix = prefix
	}
}

var defaultZapLogger = &ZapLogger{
	debug:     false,
	addCaller: false,
	days:      int64(90),

	logLevel: _defaultLevel,
}

// New init a log
func New(opts ...ZapLoggerOption) *ZapLogger {
	s := defaultZapLogger

	for _, o := range opts {
		o(s)
	}

	if len(s.path) == 0 {
		s.path = defaultLogPath(subPath)
	}

	if len(s.prefix) == 0 {
		s.prefix = os.Args[0]
	}

	var core zapcore.Core

	if s.storeInDay {
		core = dayLogger(s.path, s.prefix, s.days, s.debug)
	} else {
		core = sizeLogger(s.path, s.prefix, s.debug)
	}

	log := zap.New(core)

	if s.addCaller {
		log = log.WithOptions(zap.AddCaller())
	}
	s.Log = log
	return s
}

// func (l *ZapLogger) Sync() {
// 	_ = l.Log.Sync()

// 	if _rotateLogs != nil {

// 	}
// }

// Named  name logger
func (l *ZapLogger) Named(s string) *zap.Logger {
	return l.Log.Named(s)
}
