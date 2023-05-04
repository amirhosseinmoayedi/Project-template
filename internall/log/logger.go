package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func InitLogger(logLevel string) {
	logrus.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Fatalf("Error in initiating logger")
	}

	Logger = &logrus.Logger{
		Out:   os.Stderr,
		Level: level,
		Hooks: make(logrus.LevelHooks),
		Formatter: &logrus.TextFormatter{
			ForceColors:            true,
			DisableColors:          false,
			DisableLevelTruncation: false,
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05",
		},
	}
}
