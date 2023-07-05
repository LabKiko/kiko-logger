/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/7/3
* Time: 18:29
* Software: GoLand
 */

package zapcore

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	std "log"

	kLogger "github.com/LabKiko.kiko-logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	_oddNumberErrMsg    = "Ignored key without a value."
	_nonStringKeyErrMsg = "Ignored key-value pairs with non-string keys."
)

type kZLogger struct {
	opt           Option
	logger        *zap.Logger
	_rollingFiles []io.Writer
	atomicLevel   zap.AtomicLevel
}

func NewLogger(option ...kLogger.Option) *kZLogger {
	zLog := &kZLogger{atomicLevel: zap.NewAtomicLevelAt(zap.InfoLevel)}
	zLog.logger.WithOptions(option...)

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
	return &kZLogger{
		opt: l.opt,
		logger: l.logger.With(CopyFields(map[string]interface{}{
			"error": err,
		})...).WithOptions(zap.AddCallerSkip(0)),
		atomicLevel: l.atomicLevel,
	}
}

func (l kZLogger) Debug(args ...interface{}) {
	l.log(DebugLevel, "", args, nil)
}

func (l kZLogger) Info(args ...interface{}) {
	l.log(InfoLevel, "", args, nil)
}

func (l kZLogger) Warn(args ...interface{}) {
	l.log(WarnLevel, "", args, nil)
}
}

func (l kZLogger) Error(args ...interface{}) {
	l.log(WarnLevel, "", args, nil)
}
}

func (l kZLogger) Fatal(args ...interface{}) {
	l.log(FatalLevel, "", args, nil)
}

func (l kZLogger) Debugf(template string, args ...interface{}) {
	l.log(DebugLevel, template, args, nil)
}

func (l kZLogger) Infof(template string, args ...interface{}) {
	l.log(InfoLevel, template, args, nil)
}

func (l kZLogger) Warnf(template string, args ...interface{}) {
	l.log(WarnLevel, template, args, nil)
}

func (l kZLogger) Errorf(template string, args ...interface{}) {
	l.log(ErrorLevel, template, args, nil)
}

func (l kZLogger) Fatalf(template string, args ...interface{}) {
	l.log(FatalLevel, template, args, nil)
}

func (l kZLogger) StdLog() *log.Logger {
	stdLogger := std.New(logWriter{logFunc: func() func(msg string, fields ...interface{}) {
		logger :=&kZLogger{logger: l.logger.WithOptions(zap.AddCallerSkip(3))}
		return logger.Infof
	}},"",0)
	return stdLogger
}

func (l kZLogger) Sync() error {
	if l.logger != nil {
		return l.logger.Sync()
	}

	for _, w := range l._rollingFiles {
		r, ok := w.(*RollingFile)
		if ok {
			r.Close()
		}
	}

	return nil
}
func (l kZLogger) AddHooks(hooks ...kLogger.Hook)  {

    // zHooks := make([]func(entry zapcore.Entry)error,0,len(hooks))
	// for _,hook :=range hooks{
	//
	//
	// }

}

func CopyFields(fields map[string]interface{}) []zap.Field {
	dst := make([]zap.Field, 0, len(fields))
	for l, v := range fields {
		dst = append(dst, zap.Any(l, v))
	}
	return dst
}

func (l kZLogger) log(level Level, template string, fmArgs []interface{}, context []interface{}) {
	if level < DebugLevel || !l.logger.Core().Enabled(level.unmarshalZapLevel()) {
		return
	}
	msg := getMessage(template, fmArgs)
	if ce := l.logger.Check(zapcore.Level(level), msg); ce != nil {
		ce.Write(l.sweetenFields(context)...)
	}
}


type logWriter struct {
	logFunc func() func(msg string, fields ...interface{})
}

func (l logWriter) Write(p []byte) (int, error) {
	p = bytes.TrimSpace(p)
	if l.logFunc != nil {
		l.logFunc()(string(p))
	}
	return len(p), nil
}


// getMessage format with Sprint, Sprintf, or neither.
func getMessage(template string, fmtArgs []interface{}) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}

	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}

func (l kZLogger) sweetenFields(args []interface{}) []zap.Field {
	if len(args) == 0 {
		return nil
	}

	// Allocate enough space for the worst case; if users pass only structured
	// fields, we shouldn't penalize them with extra allocations.
	fields := make([]zap.Field, 0, len(args))
	var invalid invalidPairs

	for i := 0; i < len(args); {
		// This is a strongly-typed field. Consume it and move on.
		if f, ok := args[i].(zap.Field); ok {
			fields = append(fields, f)
			i++
			continue
		}

		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			l.logger.Error(_oddNumberErrMsg, zap.Any("ignored", args[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			// Subsequent errors are likely, so allocate once up front.
			if cap(invalid) == 0 {
				invalid = make(invalidPairs, 0, len(args)/2)
			}
			invalid = append(invalid, invalidPair{i, key, val})
		} else {
			fields = append(fields, zap.Any(keyStr, val))
		}
		i += 2
	}

	// If we encountered any invalid key-value pairs, log an error.
	if len(invalid) > 0 {
		l.logger.Error(_nonStringKeyErrMsg, zap.Array("invalid", invalid))
	}
	return fields
}

type invalidPair struct {
	position   int
	key, value interface{}
}

func (p invalidPair) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt64("position", int64(p.position))
	zap.Any("key", p.key).AddTo(enc)
	zap.Any("value", p.value).AddTo(enc)
	return nil
}

type invalidPairs []invalidPair
