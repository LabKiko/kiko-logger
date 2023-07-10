/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/7/3
* Time: 14:28
* Software: GoLand
 */

package kiko_logger

import (
	"context"
	"log"
)

type Logger interface {
	// Options The Logger options
	// Options() Option

	// SetLevel set logger level
	SetLevel(lv Level)

	// WithContext with context
	WithContext(ctx context.Context) Logger

	// WithFields set fields to always be logged
	WithFields(fields map[string]interface{}) Logger

	// WithCallDepth  with logger call depth.
	WithCallDepth(callDepth int) Logger

	// WithError with logger error
	WithError(err error) Logger

	// Debug uses fmt.Sprint to construct and log a message.
	Debug(args ...interface{})

	// Info uses fmt.Sprint to construct and log a message.
	Info(args ...interface{})

	// Warn uses fmt.Sprint to construct and log a message.
	Warn(args ...interface{})

	// Error uses fmt.Sprint to construct and log a message.
	Error(args ...interface{})

	// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
	Fatal(args ...interface{})

	// Debugf uses fmt.Sprintf to log a templated message.
	Debugf(template string, args ...interface{})

	// Infof uses fmt.Sprintf to log a templated message.
	Infof(template string, args ...interface{})

	// Warnf uses fmt.Sprintf to log a templated message.
	Warnf(template string, args ...interface{})

	// Errorf uses fmt.Sprintf to log a templated message.
	Errorf(template string, args ...interface{})

	// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
	Fatalf(template string, args ...interface{})

	// Log writes a log entry
	// Log(level Level, template string, fmtArgs []interface{}, context []interface{})

	StdLog() *log.Logger

	// Sync logger sync
	Sync() error
	// AddHooks Hooks registers functions which will be called each time the Logger writes
	// out an Entry. Repeated use of Hooks is additive.
	//
	// Hooks are useful for simple side effects, like capturing metrics for the
	// number of emitted logs. More complex side effects, including anything that
	// requires access to the Entry's structured fields, should be implemented as
	// a zapcore.Core instead. See zapcore.RegisterHooks for details.
	// AddHooks(hooks ...Hook)
}
