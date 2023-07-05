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
	"io"
	"log"

	kLogger "github.com/LabKiko.kiko-logger"
	"github.com/sirupsen/logrus"
)

type KLLogger struct {
	base      *logrus.Entry
	out       io.Writer
	formatter logrus.Formatter
}

func (l KLLogger) Options() kLogger.Option {
	// TODO implement me
	panic("implement me")
}

func (l KLLogger) SetLevel(lv kLogger.Level) {
	// TODO implement me
	panic("implement me")
}

func (l KLLogger) WithContext(ctx context.Context) kLogger.Logger {
	spanId := kLogger.ExtractSpanId(ctx)
	traceId := kLogger.ExtractTraceId(ctx)

	fields := logrus.Fields{}
	if len(spanId) > 0 {
		fields[kLogger.SpanKey] = spanId
	}
	if len(traceId) > 0 {
		fields[kLogger.TraceId] = traceId
	}

	return &KLLogger{
		base:      l.base.WithFields(fields),
		out:       l.out,
		formatter: l.formatter,
	}

}

func (l KLLogger) WithFields(fields map[string]interface{}) kLogger.Logger {
	return &KLLogger{
		base:      l.base.WithFields(fields),
		out:       l.out,
		formatter: l.formatter,
	}
}

func (l KLLogger) WithCallDepth(callDepth int) kLogger.Logger {
	l.base.Logger.SetReportCaller(true)
	return &KLLogger{
		base:      l.base,
		out:       l.out,
		formatter: l.formatter,
	}
}

func (l KLLogger) WithError(err error) kLogger.Logger {
	return &KLLogger{
		base:      l.base.WithError(err),
		out:       l.out,
		formatter: l.formatter,
	}
}

func (l KLLogger) Debug(args ...interface{}) {
	l.base.Debug(args)
}

func (l KLLogger) Info(args ...interface{}) {
	l.base.Info(args)
}

func (l KLLogger) Warn(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (l KLLogger) Error(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (l KLLogger) Fatal(args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (l KLLogger) Debugf(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (l KLLogger) Infof(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (l KLLogger) Warnf(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (l KLLogger) Errorf(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (l KLLogger) Fatalf(template string, args ...interface{}) {
	// TODO implement me
	panic("implement me")
}

func (l KLLogger) StdLog() *log.Logger {
	// TODO implement me
	panic("implement me")
}

func (l KLLogger) Sync() error {
	// TODO implement me
	panic("implement me")
}

func (l KLLogger) AddHooks(hooks ...kLogger.Hook) {
	// TODO implement me
	panic("implement me")
}

func NewLogger() kLogger.Logger {
	return &KLLogger{}
}
