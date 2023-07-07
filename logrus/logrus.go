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
	base *logrus.Logger
	opt  *Option
}

func NewLogger(opts ...Options) kLogger.Logger {
	logger := &KLLogger{
		base: logrus.New(),
	}

	logger.withOption(opts...)

	return logger
}

func (l *KLLogger) withOption(opts ...Options) {
	opt := defaultJsonOption
	for _, option := range opts {
		option(&opt)
	}

	l.base.SetLevel(opt.level)
	l.base.SetReportCaller(opt.reportCaller)

	l.base.SetFormatter(opt.formatter)
	l.base.SetBufferPool(opt.bufferPoll)

	for _, hook := range opt.hooks {
		l.base.AddHook(hook)
	}
	l.opt = &opt

}
func (l *KLLogger) Option() *Option {
	return l.opt
}

func (l *KLLogger) SetLevel(lv kLogger.Level) {
	l.base.SetLevel(convertLevel(lv))
}

func convertLevel(level kLogger.Level) logrus.Level {
	defaultLevel := logrus.InfoLevel
	if lLevel, ok := LevelMap[level]; ok {
		defaultLevel = lLevel
	}

	return defaultLevel
}

func (l *KLLogger) WithContext(ctx context.Context) kLogger.Logger {
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
		base: l.base.WithFields(fields).WithContext(ctx).Logger,
	}

}

func (l *KLLogger) WithFields(fields map[string]interface{}) kLogger.Logger {
	return &KLLogger{
		base: l.base.WithFields(fields).Logger,
	}
}

func (l *KLLogger) WithCallDepth(callDepth int) kLogger.Logger {
	l.base.SetReportCaller(true)
	return &KLLogger{
		base: l.base,
	}
}

func (l *KLLogger) WithError(err error) kLogger.Logger {
	return &KLLogger{
		base: l.base.WithError(err).Logger,
	}
}

func (l *KLLogger) Debug(args ...interface{}) {
	l.base.Debug(args...)
}

func (l *KLLogger) Info(args ...interface{}) {
	l.base.Info(args...)
}

func (l *KLLogger) Warn(args ...interface{}) {
	l.base.Warn(args...)
}

func (l *KLLogger) Error(args ...interface{}) {
	l.base.Error(args...)
}

func (l *KLLogger) Fatal(args ...interface{}) {
	l.base.Fatal(args...)
}

func (l *KLLogger) Debugf(template string, args ...interface{}) {
	l.base.Debugf(template, args...)
}

func (l *KLLogger) Infof(template string, args ...interface{}) {
	l.base.Infof(template, args...)
}

func (l *KLLogger) Warnf(template string, args ...interface{}) {
	l.base.Warnf(template, args...)
}

func (l *KLLogger) Errorf(template string, args ...interface{}) {
	l.base.Errorf(template, args...)
}

func (l *KLLogger) Fatalf(template string, args ...interface{}) {
	l.base.Fatalf(template, args...)
}

func (l *KLLogger) StdLog() *log.Logger {

	stdLogger := log.New(kLogger.LogWriter{
		LogFunc: func() func(msg string, args ...interface{}) {
			logger := &KLLogger{base: l.base}
			return logger.Infof
		},
	}, "", 0)
	return stdLogger

	return nil
}

func (l *KLLogger) Sync() error {
	// 数据落盘
	l.base.SetOutput(io.MultiWriter(l.opt.outPut...))
	return nil
}

func (l *KLLogger) AddHooks(hooks ...logrus.Hook) {
	for _, hook := range hooks {
		l.base.AddHook(hook)
	}
}
