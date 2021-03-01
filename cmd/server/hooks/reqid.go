package hooks

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"

	"moocss.com/gaea/pkg/ctxkit"
	"moocss.com/gaea/pkg/trace"
	"moocss.com/gaea/pkg/twirp"
)

// NewRequestID 生成唯一请求标识并记录到 ctx
func NewRequestID() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			ctx = context.WithValue(ctx, ctxkit.StartTimeKey, time.Now())

			traceID := trace.GetTraceID(ctx)
			twirp.SetHTTPResponseHeader(ctx, "x-trace-id", traceID)

			ctx = ctxkit.WithTraceID(ctx, traceID)

			return ctx, nil
		},
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			pkg, _ := twirp.PackageName(ctx)
			service, _ := twirp.ServiceName(ctx)
			method, _ := twirp.MethodName(ctx)

			api := "/" + pkg + "." + service + "/" + method

			span, ctx := opentracing.StartSpanFromContext(ctx, api)
			ctx = context.WithValue(ctx, spanKey, span)

			return ctx, nil
		},
		ResponseSent: func(ctx context.Context) {
			if span, ok := ctx.Value(spanKey).(opentracing.Span); ok {
				span.Finish()
			}
		},
	}
}
