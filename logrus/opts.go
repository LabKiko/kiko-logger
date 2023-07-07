/**
* Author: Jason
* Email:  jason_w96@163.com
* Date: 2023/7/3
* Time: 18:48
* Software: GoLand
 */

package logrus

import (
	"fmt"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type Options func(o *Option)

type Option struct {
	fileName     string
	level        logrus.Level
	outPut       []io.Writer
	formatter    logrus.Formatter
	hooks        []logrus.Hook
	dir          string
	timeFormat   string
	stdout       bool
	reportCaller bool
	bufferPoll   logrus.BufferPool
}

func WithStdOut(stdout bool) Options {
	return func(o *Option) {
		o.stdout = stdout
	}
}

func WithLevel(level Level) Options {
	return func(o *Option) {
		o.level = logrus.Level(level)
	}
}

func WithFileName(fileName string) Options {
	return func(o *Option) {
		o.fileName = fileName
	}
}

func WithOutPut(output ...io.Writer) Options {
	return func(o *Option) {
		o.outPut = append(o.outPut, output...)
	}
}

func WithHook(hook ...logrus.Hook) Options {
	return func(o *Option) {
		o.hooks = append(o.hooks, hook...)
	}
}

func WithFormatter(formatter logrus.Formatter) Options {
	return func(o *Option) {
		o.formatter = formatter
	}
}

func WithTimeFormat(timeStr string) Options {
	return func(o *Option) {
		o.timeFormat = timeStr
	}
}

// WithRotateLog 处理数据
func WithRotateLog(rotateOpts ...rotatelogs.Option) Options {

	return func(o *Option) {
		logfileName := fmt.Sprintf("%s/%%Y%%m%%d.log", o.dir)
		if len(rotateOpts) > 0 {
			rotateLog, _ := rotatelogs.New(
				logfileName,
				rotateOpts...,
			)
			o.outPut = append(o.outPut, rotateLog)
		}

	}
}

var (
	defaultTimeFormatter = time.DateTime
	defaultJSONFormatter = logrus.JSONFormatter{TimestampFormat: defaultTimeFormatter}
	defaultTextFormatter = logrus.TextFormatter{TimestampFormat: defaultTimeFormatter}
	// 默认json 格式

	defaultJsonOption = Option{
		level:        logrus.InfoLevel,
		outPut:       []io.Writer{os.Stdout},
		formatter:    &defaultJSONFormatter,
		dir:          "log",
		timeFormat:   defaultTimeFormatter,
		reportCaller: true,
	}

	defaultTextOption = Option{
		level:        logrus.InfoLevel,
		outPut:       []io.Writer{os.Stdout},
		formatter:    &defaultTextFormatter,
		dir:          "log",
		timeFormat:   defaultTimeFormatter,
		reportCaller: true,
	}
)
