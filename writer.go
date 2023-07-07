/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/7/7
* Time: 18:42
* Software: GoLand
 */

package kiko_logger

import (
	"bytes"
)

type LogWriter struct {
	LogFunc func() func(msg string, fields ...interface{})
}

func (l LogWriter) Write(p []byte) (int, error) {
	p = bytes.TrimSpace(p)
	if l.LogFunc != nil {
		l.LogFunc()(string(p))
	}
	return len(p), nil
}
