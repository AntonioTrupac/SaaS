package logger

import (
	"log"
	"os"
	"time"

	dbLog "gorm.io/gorm/logger"
)

func LogGORM() dbLog.Interface {
	newLogger := dbLog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		dbLog.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  dbLog.Info,  // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)
	return newLogger
}
