package logger

import (
	"log"

	"github.com/sirupsen/logrus"
)

func NewLogger() *log.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{}

	return log.New(logger.Writer(), "gorpo: ", log.Ldate|log.Ltime)
}
