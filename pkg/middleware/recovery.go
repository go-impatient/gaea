package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"moocss.com/gaea/pkg/ctxkit"
	"moocss.com/gaea/pkg/log"
	"moocss.com/gaea/pkg/trace"
)

func startSpan(r *http.Request) (*http.Request, opentracing.Span) {
	operation := "ServerHTTP"

	ctx := r.Context()
	var span opentracing.Span

	tracer := opentracing.GlobalTracer()
	carrier := opentracing.HTTPHeadersCarrier(r.Header)

	if spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, carrier); err == nil {
		span = opentracing.StartSpan(operation, ext.RPCServerOption(spanCtx))
		ctx = opentracing.ContextWithSpan(ctx, span)
	} else {
		span, ctx = opentracing.StartSpanFromContext(ctx, operation)
	}

	ext.SpanKindRPCServer.Set(span)
	span.SetTag(string(ext.HTTPUrl), r.URL.Path)

	return r.WithContext(ctx), span
}

// Recovery is a server middleware that recovers from any panics.
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r, span := startSpan(r)
		defer func() {
			if rec := recover(); rec != nil {
				ctx := r.Context()
				ctx = ctxkit.WithTraceID(ctx, trace.GetTraceID(ctx))
				log.Get(ctx).Error(rec, string(debug.Stack()))
			}
			span.Finish()
		}()

		next.ServeHTTP(w, r)
	})
}
