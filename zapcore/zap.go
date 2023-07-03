/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/7/3
* Time: 18:29
* Software: GoLand
 */

package zapcore

import (
	"context"
	"log"

	kLogger "github.com/LabKiko.kiko-logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type kZLogger struct {
	opt    Option
	logger *zap.Logger
}

func NewKZLogger(zapCore zapcore.Core, zapOption zap.Option) kLogger.Logger {
	return &kZLogger{logger: zap.New(zapCore, zapOption)}
}

func (k kZLogger) Options() kLogger.Option {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) SetLevel(lv kLogger.Level) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) WithContext(ctx context.Context) kLogger.Logger {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) WithFields(fields map[string]interface{}) kLogger.Logger {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) WithCallDepth(callDepth int) kLogger.Logger {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) WithError(err error) kLogger.Logger {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Debug(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Info(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Warn(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Error(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Fatal(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Debugf(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Infof(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Warnf(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Errorf(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Fatalf(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Debugw(msg string, keysAndValues ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Infow(msg string, keysAndValues ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Warnw(msg string, keysAndValues ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Errorw(msg string, keysAndValues ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Log(level kLogger.Level, template string, fmtArgs []interface{}, context []interface{}) {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) StdLog() *log.Logger {
	// TODO implement me
	panic("implement me")
}

func (k kZLogger) Sync() error {
	// TODO implement me
	panic("implement me")
}
