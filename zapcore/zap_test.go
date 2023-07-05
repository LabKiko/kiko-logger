/**
* Author: Jason
* Email: wudongdong@rongma.com
* Date: 2023/7/5
* Time: 18:35
* Software: GoLand
 */

package zapcore

import (
	"context"
	"io"
	std "log"
	"reflect"
	"testing"

	kiko_logger "github.com/LabKiko.kiko-logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestCopyFields(t *testing.T) {
	type args struct {
		fields map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []zap.Field
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyFields(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_AddHooks(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		hooks []kiko_logger.Hook
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.AddHooks(tt.args.hooks...)
		})
	}
}

func TestKzLog_Debug(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Debug(tt.args.args...)
		})
	}
}

func TestKzLog_Debugf(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Debugf(tt.args.template, tt.args.args...)
		})
	}
}

func TestKzLog_Debugw(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Debugw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestKzLog_Error(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Error(tt.args.args...)
		})
	}
}

func TestKzLog_Errorf(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Errorf(tt.args.template, tt.args.args...)
		})
	}
}

func TestKzLog_Errorw(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Errorw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestKzLog_Fatal(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Fatal(tt.args.args...)
		})
	}
}

func TestKzLog_Fatalf(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Fatalf(tt.args.template, tt.args.args...)
		})
	}
}

func TestKzLog_Fatalw(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Fatalw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestKzLog_Info(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Info(tt.args.args...)
		})
	}
}

func TestKzLog_Infof(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Infof(tt.args.template, tt.args.args...)
		})
	}
}

func TestKzLog_Infow(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Infow(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestKzLog_LevelEnablerFunc(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		level zapcore.Level
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   LevelEnablerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if got := l.LevelEnablerFunc(tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelEnablerFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_Log(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		level    Level
		template string
		fmtArgs  []interface{}
		context  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Log(tt.args.level, tt.args.template, tt.args.fmtArgs, tt.args.context)
		})
	}
}

func TestKzLog_Options(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	tests := []struct {
		name   string
		fields fields
		want   kiko_logger.Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if got := l.Options(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Options() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_SetLevel(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		lv kiko_logger.Level
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.SetLevel(tt.args.lv)
		})
	}
}

func TestKzLog_StdLog(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	tests := []struct {
		name   string
		fields fields
		want   *std.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if got := l.StdLog(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StdLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_Sync(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if err := l.Sync(); (err != nil) != tt.wantErr {
				t.Errorf("Sync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKzLog_Warn(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Warn(tt.args.args...)
		})
	}
}

func TestKzLog_Warnf(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Warnf(tt.args.template, tt.args.args...)
		})
	}
}

func TestKzLog_Warnw(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.Warnw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestKzLog_WithCallDepth(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		callDepth int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   kiko_logger.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if got := l.WithCallDepth(tt.args.callDepth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCallDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_WithContext(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   kiko_logger.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if got := l.WithContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_WithError(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   kiko_logger.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if got := l.WithError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_WithFields(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		fields map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   kiko_logger.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if got := l.WithFields(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_clone(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	tests := []struct {
		name   string
		fields fields
		want   *KzLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if got := l.clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_createOutput(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    zapcore.WriteSyncer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			got, err := l.createOutput(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("createOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createOutput() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_extractSpanId(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if got := l.extractSpanId(tt.args.ctx); got != tt.want {
				t.Errorf("extractSpanId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_extractTraceId(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if got := l.extractTraceId(tt.args.ctx); got != tt.want {
				t.Errorf("extractTraceId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_log(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		level    Level
		template string
		fmtArgs  []interface{}
		context  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.log(tt.args.level, tt.args.template, tt.args.fmtArgs, tt.args.context)
		})
	}
}

func TestKzLog_setUp(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if err := l.setUp(); (err != nil) != tt.wantErr {
				t.Errorf("setUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKzLog_setupWithConsole(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	tests := []struct {
		name    string
		fields  fields
		want    []zapcore.Core
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			got, err := l.setupWithConsole()
			if (err != nil) != tt.wantErr {
				t.Errorf("setupWithConsole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupWithConsole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_setupWithFile(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	tests := []struct {
		name    string
		fields  fields
		want    []zapcore.Core
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			got, err := l.setupWithFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("setupWithFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupWithFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_setupWithFileConsole(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	tests := []struct {
		name    string
		fields  fields
		want    []zapcore.Core
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			got, err := l.setupWithFileConsole()
			if (err != nil) != tt.wantErr {
				t.Errorf("setupWithFileConsole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupWithFileConsole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_setupWithFiles(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	tests := []struct {
		name    string
		fields  fields
		want    []zapcore.Core
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			got, err := l.setupWithFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("setupWithFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupWithFiles() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_sweetenFields(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []zap.Field
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			if got := l.sweetenFields(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sweetenFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKzLog_withOptions(t *testing.T) {
	type fields struct {
		opt           Option
		base          *zap.Logger
		_rollingFiles []io.Writer
		atomicLevel   zap.AtomicLevel
	}
	type args struct {
		opts []Options
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KzLog{
				opt:           tt.fields.opt,
				base:          tt.fields.base,
				_rollingFiles: tt.fields._rollingFiles,
				atomicLevel:   tt.fields.atomicLevel,
			}
			l.withOptions(tt.args.opts...)
		})
	}
}

func TestMust(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Must(tt.args.err)
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		opts []Options
	}
	tests := []struct {
		name string
		args args
		want *KzLog
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMessage(t *testing.T) {
	type args struct {
		template string
		fmtArgs  []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMessage(tt.args.template, tt.args.fmtArgs); got != tt.want {
				t.Errorf("getMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_invalidPair_MarshalLogObject(t *testing.T) {
	type fields struct {
		position int
		key      interface{}
		value    interface{}
	}
	type args struct {
		enc zapcore.ObjectEncoder
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := invalidPair{
				position: tt.fields.position,
				key:      tt.fields.key,
				value:    tt.fields.value,
			}
			if err := p.MarshalLogObject(tt.args.enc); (err != nil) != tt.wantErr {
				t.Errorf("MarshalLogObject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_invalidPairs_MarshalLogArray(t *testing.T) {
	type args struct {
		enc zapcore.ArrayEncoder
	}
	tests := []struct {
		name    string
		ps      invalidPairs
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ps.MarshalLogArray(tt.args.enc); (err != nil) != tt.wantErr {
				t.Errorf("MarshalLogArray() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_logWriter_Write(t *testing.T) {
	type fields struct {
		logFunc func() func(msg string, fields ...interface{})
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logWriter{
				logFunc: tt.fields.logFunc,
			}
			got, err := l.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Write() got = %v, want %v", got, tt.want)
			}
		})
	}
}
