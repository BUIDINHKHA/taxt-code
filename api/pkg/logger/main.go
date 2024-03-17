package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func InitLog() *Logger {
	format := logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		DisableQuote:    true,
		DisableColors:   true,
		FieldMap: logrus.FieldMap{
			"level": "logLevel",
		},
	}
	var log = logrus.New()
	log.SetFormatter(&format)
	return &Logger{
		logger: log,
	}
}
func (l *Logger) Info(id, message string) {
	l.logger.WithFields(logrus.Fields{
		"id": id,
	}).Info(message)
}

func (l *Logger) Error(id, message, err string) {
	l.logger.WithFields(logrus.Fields{
		"id":    id,
		"error": err,
	}).Error(message)
}

func (l *Logger) Fatal(id, message, err string) {
	l.logger.WithFields(logrus.Fields{
		"id":    id,
		"error": err,
	}).Fatal(message)
}
