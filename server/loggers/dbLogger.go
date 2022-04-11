package logger

import (
	"io"
	"log"
	"os"
	"time"

	dbLog "gorm.io/gorm/logger"
)

func LogGORM() dbLog.Interface {
	f, _ := os.Create("gorm.log")

	newLogger := dbLog.New(
		log.New(io.MultiWriter(f), "\r\n", log.LstdFlags), // io writer
		dbLog.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  dbLog.Info,  // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)

	return newLogger
}
