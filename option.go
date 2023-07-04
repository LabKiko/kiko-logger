/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/7/3
* Time: 18:26
* Software: GoLand
 */

package kiko_logger

import (
	"go.uber.org/zap/zapcore"
)

type Encoder string
type Option struct {
	// The logging level the Logger should log at. default is `InfoLevel`
	level Level
	// basePath defines base path of log file
	basePath string
	//  Logger file name
	filename string
	// enableConsole display logs to standard output
	console bool
	// disableDisk disable rolling file
	disableDisk bool

	callerSkip int

	namespace string

	fields map[string]interface{}

	encoder Encoder

	encoderConfig zapcore.EncoderConfig
}
