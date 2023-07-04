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
	"io"
	"log"

	kLogger "github.com/LabKiko.kiko-logger"
	"go.uber.org/zap"
)

type kZLogger struct {
	opt           Option
	logger        *zap.Logger
	_rollingFiles []io.Writer
	atomicLevel   zap.AtomicLevel
}

func NewLogger(option ...kLogger.Option) *kZLogger {
	zLog := &kZLogger{atomicLevel: zap.NewAtomicLevelAt(zap.InfoLevel)}
	zLog.logger.WithOptions()

	return zLog
}

func (l kZLogger) Options() kLogger.Option {
	return kLogger.Option{}
}

func (l kZLogger) SetLevel(lv kLogger.Level) {
	l.opt.level = Level(lv)
	l.atomicLevel.SetLevel(l.opt.level.unmarshalZapLevel())
}

func (l kZLogger) WithContext(ctx context.Context) kLogger.Logger {
	spanId := kLogger.ExtractSpanId(ctx)
	traceId := kLogger.ExtractTraceId(ctx)

	fields := make(map[string]interface{})
	if len(spanId) > 0 {
		fields[kLogger.SpanKey] = spanId
	}
	if len(traceId) > 0 {
		fields[kLogger.TraceId] = spanId
	}

	return &kZLogger{opt: l.opt, logger: l.logger.With(CopyFields(fields)...).WithOptions(zap.AddCallerSkip(0))}
}

func (l kZLogger) WithFields(fields map[string]interface{}) kLogger.Logger {
	return &kZLogger{
		opt:    l.opt,
		logger: l.logger.With(CopyFields(fields)...).WithOptions(zap.AddCallerSkip(0)),
	}
}

func (l kZLogger) WithCallDepth(callDepth int) kLogger.Logger {
	return &kZLogger{
		opt:    l.opt,
		logger: l.logger.WithOptions(zap.AddCallerSkip(callDepth)),
	}
}

func (l kZLogger) WithError(err error) kLogger.Logger {

}

func (l kZLogger) Debug(args ...interface{}) {

}

func (l kZLogger) Info(args ...interface{}) {

}

func (l kZLogger) Warn(args ...interface{}) {

}

func (l kZLogger) Error(args ...interface{}) {

}

func (l kZLogger) Fatal(args ...interface{}) {

}

func (l kZLogger) Debugf(template string, args ...interface{}) {

}

func (l kZLogger) Infof(template string, args ...interface{}) {

}

func (l kZLogger) Warnf(template string, args ...interface{}) {

}

func (l kZLogger) Errorf(template string, args ...interface{}) {

}

func (l kZLogger) Fatalf(template string, args ...interface{}) {

}

func (l kZLogger) Log(level kLogger.Level, template string, fmtArgs []interface{}, context []interface{}) {

}

func (l kZLogger) AddHooks() kLogger.Option {

}

func (l kZLogger) StdLog() *log.Logger {

}

func (l kZLogger) Sync() error {
	return l.logger.Sync()
}

func CopyFields(fields map[string]interface{}) []zap.Field {
	dst := make([]zap.Field, 0, len(fields))
	for l, v := range fields {
		dst = append(dst, zap.Any(l, v))
	}
	return dst
}
