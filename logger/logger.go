// logger.go
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger() {
	log = logrus.New()

	log.Out = os.Stdout

	// You can set the log level, default is Info
	log.Level = logrus.InfoLevel

	// You can also format the logs in different formats like JSONFormat
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}
}

func Log() *logrus.Logger {
	return log
}
