/**
* Author: Jason
* Email: wudongdong@rongma.com
* Date: 2023/8/21
* Time: 14:45
* Software: GoLand
 */

package logrus

import (
	"context"
	"io"
	"log"
	"reflect"
	"testing"

	kiko_logger "github.com/LabKiko/kiko-logger/logger"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func TestKLLogger_AddHooks(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
	}
	type args struct {
		hooks []logrus.Hook
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.AddHooks(tt.args.hooks...)
		})
	}
}

func TestKLLogger_Debug(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.Debug(tt.args.args...)
		})
	}
}

func TestKLLogger_Debugf(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.Debugf(tt.args.template, tt.args.args...)
		})
	}
}

func TestKLLogger_Error(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.Error(tt.args.args...)
		})
	}
}

func TestKLLogger_Errorf(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.Errorf(tt.args.template, tt.args.args...)
		})
	}
}

func TestKLLogger_Fatal(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.Fatal(tt.args.args...)
		})
	}
}

func TestKLLogger_Fatalf(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.Fatalf(tt.args.template, tt.args.args...)
		})
	}
}

func TestKLLogger_Info(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.Info(tt.args.args...)
		})
	}
}

func TestKLLogger_Infof(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.Infof(tt.args.template, tt.args.args...)
		})
	}
}

func TestKLLogger_Option(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
	}
	tests := []struct {
		name   string
		fields fields
		want   *Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			if got := l.Option(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Option() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKLLogger_SetLevel(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.SetLevel(tt.args.lv)
		})
	}
}

func TestKLLogger_StdLog(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
	}
	tests := []struct {
		name   string
		fields fields
		want   *log.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			if got := l.StdLog(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StdLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKLLogger_Sync(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			if err := l.Sync(); (err != nil) != tt.wantErr {
				t.Errorf("Sync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKLLogger_Warn(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.Warn(tt.args.args...)
		})
	}
}

func TestKLLogger_Warnf(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.Warnf(tt.args.template, tt.args.args...)
		})
	}
}

func TestKLLogger_WithCallDepth(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			if got := l.WithCallDepth(tt.args.callDepth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCallDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKLLogger_WithContext(t *testing.T) {
	logger := NewLogger()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "trace_id", "dddd")
	logger.WithContext(ctx).Info("failed to call data")
}

func TestKLLogger_WithError(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			if got := l.WithError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKLLogger_WithFields(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			if got := l.WithFields(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKLLogger_withOption(t *testing.T) {
	type fields struct {
		base *logrus.Logger
		opt  *Option
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
			l := &KLLogger{
				base: tt.fields.base,
				opt:  tt.fields.opt,
			}
			l.withOption(tt.args.opts...)
		})
	}
}

func TestNewLogger(t *testing.T) {
	type args struct {
		opts []Options
	}
	tests := []struct {
		name string
		args args
		want kiko_logger.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithFileName(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithFileName(tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithFormatter(t *testing.T) {
	type args struct {
		formatter logrus.Formatter
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithFormatter(tt.args.formatter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithHook(t *testing.T) {
	type args struct {
		hook []logrus.Hook
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithHook(tt.args.hook...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLevel(t *testing.T) {
	type args struct {
		level Level
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithLevel(tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithOutPut(t *testing.T) {
	type args struct {
		output []io.Writer
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithOutPut(tt.args.output...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithOutPut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithRotateLog(t *testing.T) {
	type args struct {
		rotateOpts []rotatelogs.Option
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithRotateLog(tt.args.rotateOpts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRotateLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithStdOut(t *testing.T) {
	type args struct {
		stdout bool
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithStdOut(tt.args.stdout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithStdOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithTimeFormat(t *testing.T) {
	type args struct {
		timeStr string
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithTimeFormat(tt.args.timeStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTimeFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertLevel(t *testing.T) {
	type args struct {
		level kiko_logger.Level
	}
	tests := []struct {
		name string
		args args
		want logrus.Level
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertLevel(tt.args.level); got != tt.want {
				t.Errorf("convertLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
