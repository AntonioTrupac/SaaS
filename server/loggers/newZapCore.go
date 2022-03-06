package logger

import (
	"time"

	"github.com/spf13/afero"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var defaultLumberJack = &lumberjack.Logger{
	Filename:   "uber-zap.log",
	MaxSize:    100, // Megabytes
	MaxBackups: 31,
	MaxAge:     31, // days
	Compress:   false,
}

// newZapCore initializa a zap log
func newZapCore(path, prefix string, level zapcore.LevelEnabler) zapcore.Core {
	dateTimeFmtInFile := time.Now().Format("2006-01-02-15")

	var logPath string

	if len(path) == 0 {
		logPath = defaultLogPath(path)
	} else {
		logPath = path
	}

	afs := afero.NewOsFs()

	check, _ := afero.DirExists(afs, logPath) // path exists? and is a directory?

	if !check {
		err := afs.MkdirAll(logPath, 0755)
		if err != nil {
			panic("Can't make path" + logPath)
		}
	}

	var w zapcore.WriteSyncer
	var logFileName string

	if len(prefix) == 0 {
		prefix = "default"
	}

	logFileName = logPath + "/" + prefix + "-" + dateTimeFmtInFile + ".log"

	defaultLumberJack.Filename = logFileName

	w = zapcore.AddSync(defaultLumberJack)

	return zapcore.NewCore(zapcore.NewJSONEncoder(defaultEncoder), w, level)
}

var defaultEncoder = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "level",
	NameKey:        "name",
	CallerKey:      "caller",
	MessageKey:     "msg",
	StacktraceKey:  "stack",
	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.NanosDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}
