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
	"fmt"
	"log"

	kLogger "github.com/LabKiko.kiko-logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type KZLogger struct {
	opt    Option
	logger *zap.Logger
	// stdLog *log.Logger
}

func NewZapCoreLogger(zapCore zapcore.Core, zapOption zap.Option) *KZLogger {
	// stdLog := log.New(logger.Writer(), "", 0)
	return &KZLogger{logger: zap.New(zapCore, zapOption)}
}

func (kz *KZLogger) WithContext(ctx context.Context) kLogger.Logger {
	return kz.logger.WithFields(map[string]interface{}{})
}
func (kz *KZLogger) Options() Option {
	// Implement your logic here
	return nil
}

func (kz *KZLogger) SetLevel(lv Level) {
	// Implement your logic here
}

func (kz *KZLogger) WithContext(ctx context.Context) kLogger.Logger {
	// Implement your logic here
	return nil
}

func (kz *KZLogger) WithFields(fields map[string]interface{}) *kLogger.Logger {
	// Implement your logic here
	return kz.logger.
}

func (kz *KZLogger) WithCallDepth(callDepth int) *kLogger.Logger {
	// Implement your logic here
	return nil
}

func (kz *KZLogger) WithError(err error) *kLogger.Logger {
	// Implement your logic here
	return nil
}

func (kz *KZLogger) Debug(args ...interface{}) *kLogger.Logger {
	l.logger.Debug(fmt.Sprint(args...))
}

func (kz *KZLogger) Info(args ...interface{}) *kLogger.Logger {
	l.logger.Info(fmt.Sprint(args...))
}

func (kz *KZLogger) Warn(args ...interface{}) *kLogger.Logger {
	l.logger.Warn(fmt.Sprint(args...))
}

func (kz *KZLogger) Error(args ...interface{}) *kLogger.Logger {
	l.logger.Error(fmt.Sprint(args...))
}

func (kz *KZLogger) Fatal(args ...interface{}) *kLogger.Logger {
	l.logger.Fatal(fmt.Sprint(args...))
}

func (kz *KZLogger) Debugf(template string, args ...interface{}) *kLogger.Logger {
	l.logger.Debug(fmt.Sprintf(template, args...))
}

func (kz *KZLogger) Infof(template string, args ...interface{}) *kLogger.Logger {
	l.logger.Info(fmt.Sprintf(template, args...))
}

func (kz *KZLogger) Warnf(template string, args ...interface{}) *kLogger.Logger {
	l.logger.Warn(fmt.Sprintf(template, args...))
}

func (kz *KZLogger) Errorf(template string, args ...interface{}) *kLogger.Logger {
	l.logger.Error(fmt.Sprintf(template, args...))
}

func (kz *KZLogger) Fatalf(template string, args ...interface{}) *kLogger.Logger {
	l.logger.Fatal(fmt.Sprintf(template, args...))
}

func (kz *KZLogger) Debugw(msg string, keysAndValues ...interface{}) *kLogger.Logger {
	l.logger.Debugw(msg, keysAndValues...)
}

func (kz *KZLogger) Infow(msg string, keysAndValues ...interface{}) *kLogger.Logger {
	l.logger.Infow(msg, keysAndValues...)
}

func (kz *KZLogger) Warnw(msg string, keysAndValues ...interface{}) *kLogger.Logger {
	l.logger.Warnw(msg, keysAndValues...)
}

func (kz *KZLogger) Errorw(msg string, keysAndValues ...interface{}) *kLogger.Logger {
	l.logger.Errorw(msg, keysAndValues...)
}

func (kz *KZLogger) Fatalw(msg string, keysAndValues ...interface{}) *kLogger.Logger {
	return kz.logger.Fatalw(msg, keysAndValues...)
}

func (kz *KZLogger) Log(level Level, template string, fmtArgs []interface{}, context []interface{}) {
	// Implement your logic here
}

func (kz *KZLogger) StdLog() *log.Logger {
	// TODO
	return nil
}

func (kz *KZLogger) Sync() error {
	err := kz.logger.Sync()
	if err != nil {
		return err
	}
	return nil
}
