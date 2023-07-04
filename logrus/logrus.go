/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/7/3
* Time: 18:48s
* Software: GoLand
 */

package logrus

import (
	"context"
	"log"

	kLogger "github.com/LabKiko.kiko-logger"
)

type KLLogger struct {
}

func (K KLLogger) Options() kLogger.Option {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) SetLevel(lv kLogger.Level) {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) WithContext(ctx context.Context) kLogger.Logger {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) WithFields(fields map[string]interface{}) kLogger.Logger {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) WithCallDepth(callDepth int) kLogger.Logger {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) WithError(err error) kLogger.Logger {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) Debug(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) Info(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) Warn(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) Error(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) Fatal(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) Debugf(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) Infof(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) Warnf(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) Errorf(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) Fatalf(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) Log(level kLogger.Level, template string, fmtArgs []interface{}, context []interface{}) {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) StdLog() *log.Logger {
	// TODO implement me
	panic("implement me")
}

func (K KLLogger) Sync() error {
	// TODO implement me
	panic("implement me")
}

func NewLogger() kLogger.Logger {
	return &KLLogger{}
}
