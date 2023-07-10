/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/7/3
* Time: 18:26
* Software: GoLand
 */

package kiko_logger

type Level int8

const (
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel = iota + 1
	// InfoLevel is the default logging priority.
	// General operational entries about what's going on inside the application.
	InfoLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	ErrorLevel
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. highest level of severity.
	FatalLevel
)
