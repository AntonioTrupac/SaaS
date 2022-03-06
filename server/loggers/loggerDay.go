package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog/diode"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DayLogger  day logger
func dayLogger(path, prefix string, days int64, debug bool) zapcore.Core {

	_errorLevelPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	if len(path) == 0 {
		path = defaultLogPath(subPath)
	}

	if len(prefix) == 0 {
		prefix = os.Args[0]
	}

	allWriter := getWriter(path+"/"+prefix+".log", days)
	errorWriter := getWriter(path+"/"+prefix+"-error.log", days)

	var core zapcore.Core

	if debug {
		stdout := getStdout()
		// 最后创建具体的Logger
		core = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewConsoleEncoder(defaultEncoder), zapcore.AddSync(stdout), _defaultLevel), // zap.NewDevelopmentEncoderConfig()
			zapcore.NewCore(zapcore.NewJSONEncoder(defaultEncoder), zapcore.AddSync(allWriter), _defaultLevel),
			zapcore.NewCore(zapcore.NewJSONEncoder(defaultEncoder), zapcore.AddSync(errorWriter), _errorLevelPriority),
		)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewJSONEncoder(defaultEncoder), zapcore.AddSync(allWriter), _defaultLevel),
			zapcore.NewCore(zapcore.NewJSONEncoder(defaultEncoder), zapcore.AddSync(errorWriter), _errorLevelPriority),
		)
	}

	return core // zap.AddCaller())
}

const (
	_interval      = 0 // time.Duration(15) * time.Microsecond
	_intervalWrite = time.Duration(15) * time.Minute
	_bufferSize    = 1024 * 4
)

var _rotatelogs *rotatelogs.RotateLogs

func getWriter(filename string, days int64) io.Writer {

	var err error
	_rotatelogs, err = rotatelogs.New(
		filename+"-%Y%m%d%H"+".json",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*time.Duration(24)*time.Duration(days)),
		rotatelogs.WithRotationTime(_intervalWrite),
	)
	if err != nil {
		return diode.NewWriter(os.Stdout, _bufferSize, _interval, func(missed int) {
			fmt.Printf("Dropped %d messages\n", missed)
		})
	}

	return diode.NewWriter(_rotatelogs, _bufferSize, _interval, func(missed int) {
		fmt.Printf("Dropped %d messages\n", missed)
	})
}

func getStdout() io.Writer {
	w := diode.NewWriter(os.Stdout, _bufferSize, _interval, func(missed int) {
		fmt.Printf("Dropped %d messages\n", missed)
	})
	return w
}
