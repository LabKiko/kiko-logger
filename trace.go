/**
* Author: Jason
* Email: wudongdong@rongma.com
* Date: 2023/7/4
* Time: 14:39
* Software: GoLand
 */

package kiko_logger

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

const (
	SpanKey   = "span_id"
	TraceId   = "trace_id"
	RequestId = "request_id"
)

func ExtractSpanId(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasSpanID() {
		return spanCtx.SpanID().String()
	}

	return ""
}

func ExtractTraceId(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		return spanCtx.TraceID().String()
	}

	return ""
}
