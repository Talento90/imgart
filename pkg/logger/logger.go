package logger

import (
	"log"

	"github.com/sirupsen/logrus"
)

type Logger struct {
}

func NewLogger() *log.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	return log.New(logger.Writer(), "gorpo: ", log.Ldate|log.Ltime)
}
