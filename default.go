package kiko_logger

import (
	"context"

	logger "github.com/LabKiko/kiko-logger/logger"
	Klogger "github.com/LabKiko/kiko-logger/logrus"
)

var defaultLogger = Klogger.NewLogger()

// WithContext with context
func WithContext(ctx context.Context) logger.Logger {
	return defaultLogger.WithContext(ctx)
}

// WithFields set fields to always be logged
func WithFields(fields map[string]interface{}) logger.Logger {
	return defaultLogger.WithFields(fields)
}

// WithCallDepth  with logger call depth.
func WithCallDepth(callDepth int) logger.Logger {
	return defaultLogger.WithCallDepth(callDepth)
}

// WithError with logger error
func WithError(err error) logger.Logger {
	return defaultLogger.WithError(err)
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	defaultLogger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	defaultLogger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
}
