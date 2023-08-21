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

	"github.com/LabKiko/kiko-logger/logger"
	utils "github.com/LabKiko/kiko-logger/utils"
	"github.com/sirupsen/logrus"
)

type KLLogger struct {
	base     *logrus.Logger
	logEntry *logrus.Entry
	opt      *Option
}

func NewLogger(opts ...Options) kiko_logger.Logger {
	logger := &KLLogger{
		base: logrus.New(),
	}

	logger.withOption(opts...)

	logger.logEntry = logrus.NewEntry(logger.base)
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

func (l *KLLogger) SetLevel(lv kiko_logger.Level) {
	l.base.SetLevel(convertLevel(lv))
}

func convertLevel(level kiko_logger.Level) logrus.Level {
	defaultLevel := logrus.InfoLevel
	if lLevel, ok := LevelMap[level]; ok {
		defaultLevel = lLevel
	}

	return defaultLevel
}

func (l *KLLogger) WithContext(ctx context.Context) kiko_logger.Logger {
	spanId := utils.ExtractSpanId(ctx)
	traceId := utils.ExtractTraceId(ctx)

	fields := logrus.Fields{}
	if len(spanId) > 0 {
		fields[utils.SpanKey] = spanId
	}
	if len(traceId) > 0 {
		fields[utils.TraceId] = traceId
	}

	return &KLLogger{
		logEntry: l.logEntry.WithFields(fields),
	}

}

func (l *KLLogger) WithFields(fields map[string]interface{}) kiko_logger.Logger {
	return &KLLogger{
		logEntry: l.logEntry.WithFields(fields),
	}
}

func (l *KLLogger) WithCallDepth(callDepth int) kiko_logger.Logger {
	l.base.SetReportCaller(true)
	return &KLLogger{
		base:     l.base,
		logEntry: l.logEntry,
	}
}

func (l *KLLogger) WithError(err error) kiko_logger.Logger {
	return &KLLogger{
		logEntry: l.logEntry.WithError(err),
	}
}

func (l *KLLogger) Debug(args ...interface{}) {
	l.logEntry.Debug(args...)
}

func (l *KLLogger) Info(args ...interface{}) {
	l.logEntry.Info(args...)
}

func (l *KLLogger) Warn(args ...interface{}) {
	l.logEntry.Warn(args...)
}

func (l *KLLogger) Error(args ...interface{}) {
	l.logEntry.Error(args...)
}

func (l *KLLogger) Fatal(args ...interface{}) {
	l.base.Fatal(args...)
}

func (l *KLLogger) Debugf(template string, args ...interface{}) {
	l.logEntry.Debugf(template, args...)
}

func (l *KLLogger) Infof(template string, args ...interface{}) {
	l.logEntry.Infof(template, args...)
}

func (l *KLLogger) Warnf(template string, args ...interface{}) {
	l.logEntry.Warnf(template, args...)
}

func (l *KLLogger) Errorf(template string, args ...interface{}) {
	l.logEntry.Errorf(template, args...)
}

func (l *KLLogger) Fatalf(template string, args ...interface{}) {
	l.logEntry.Fatalf(template, args...)
}

func (l *KLLogger) StdLog() *log.Logger {

	stdLogger := log.New(kiko_logger.LogWriter{
		LogFunc: func() func(msg string, args ...interface{}) {
			logger := &KLLogger{base: l.base}
			return logger.Infof
		},
	}, "", 0)
	return stdLogger

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
