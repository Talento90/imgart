package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger interface for logging messages
type Logger interface {
	Debug(args ...interface{})
	DebugWithFields(fields map[string]interface{}, args ...interface{})

	Info(args ...interface{})
	InfoWithFields(fields map[string]interface{}, args ...interface{})

	Warn(args ...interface{})
	WarnWithFields(fields map[string]interface{}, args ...interface{})

	Error(args ...interface{})
	ErrorWithFields(fields map[string]interface{}, args ...interface{})

	Fatal(args ...interface{})
	FatalWithFields(fields map[string]interface{}, args ...interface{})

	Panic(args ...interface{})
	PanicWithFields(fields map[string]interface{}, args ...interface{})
}

// NewLogger creates a new logger
func NewLogger(config Configuration) (Logger, error) {
	if config.Level == "" {
		config.Level = "debug"
	}

	if config.Output == nil {
		config.Output = os.Stdout
	}

	level, err := logrus.ParseLevel(config.Level)

	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.SetLevel(level)
	logger.Out = config.Output
	logger.Formatter = &logrus.TextFormatter{}

	return &logrusWrapper{Logger: logger}, nil
}

// Fields for structuring logging
type Fields map[string]interface{}

// logrusWrapper basic logger wrapper around logrus
type logrusWrapper struct {
	*logrus.Logger
}

func (l *logrusWrapper) Debug(args ...interface{}) {
	l.Logger.Debug(args)
}

func (l *logrusWrapper) DebugWithFields(fields map[string]interface{}, args ...interface{}) {
	l.Logger.WithFields(fields).Debug(args)
}

func (l *logrusWrapper) Info(args ...interface{}) {
	l.Logger.Info(args)
}

func (l *logrusWrapper) InfoWithFields(fields map[string]interface{}, args ...interface{}) {
	l.Logger.WithFields(fields).Info(args)
}

func (l *logrusWrapper) Warn(args ...interface{}) {
	l.Logger.Warn(args)
}

func (l *logrusWrapper) WarnWithFields(fields map[string]interface{}, args ...interface{}) {
	l.Logger.WithFields(fields).Warn(args)
}

func (l *logrusWrapper) Error(args ...interface{}) {
	l.Logger.Error(args)
}

func (l *logrusWrapper) ErrorWithFields(fields map[string]interface{}, args ...interface{}) {
	l.Logger.WithFields(fields).Error(args)
}

func (l *logrusWrapper) Fatal(args ...interface{}) {
	l.Logger.Fatal(args)
}

func (l *logrusWrapper) FatalWithFields(fields map[string]interface{}, args ...interface{}) {
	l.Logger.WithFields(fields).Fatal(args)
}

func (l *logrusWrapper) Panic(args ...interface{}) {
	l.Logger.Panic(args)
}

func (l *logrusWrapper) PanicWithFields(fields map[string]interface{}, args ...interface{}) {
	l.Logger.WithFields(fields).Panic(args)
}
