package config

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

func NewLogger(config *Config) *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(config.Logrus.Level))
	log.SetFormatter(&logrus.JSONFormatter{})

	logDir := "./logs"
	logFile := logDir + "/logfile.log"
	log.SetOutput(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    15,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
	})
	return log

}
