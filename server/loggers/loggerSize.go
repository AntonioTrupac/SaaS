package logger

import "go.uber.org/zap/zapcore"

func sizeLogger(path string, prefix string, debug bool) zapcore.Core {
	if debug {
		return zapcore.NewTee(zapcore.NewCore(zapcore.NewConsoleEncoder(defaultEncoder), zapcore.AddSync(getStdout()), _defaultLevel),
			newZapCore(path, prefix, _defaultLevel))
	}
	return newZapCore(path, prefix, _defaultLevel)
}
