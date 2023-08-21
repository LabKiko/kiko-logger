/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/7/3
* Time: 18:58
* Software: GoLand
 */

package kiko_logger

import (
	"context"
)

type loggerKey struct{}

func FromContext(ctx context.Context) (Logger, bool) {
	l, ok := ctx.Value(loggerKey{}).(Logger)
	return l, ok
}

func NewContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, l)
}
