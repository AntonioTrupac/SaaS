package logger

import (
	// "log"

	"go.uber.org/zap"
)

var logger *zap.Logger

// func InitLogger() {
// 	logger, err := zap.NewProduction()

// 	if err != nil {
// 		log.Fatalf("can't initialize zap logger: %v", err)
// 	}

// 	defer logger.Sync() // flushes buffer, if any
// }

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	logger.Fatal(message, fields...)
}
