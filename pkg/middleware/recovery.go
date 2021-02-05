package middleware

import (
	"context"
	"net/http"
	"runtime/debug"

	"moocss.com/gaea/pkg/ctxkit"
	"moocss.com/gaea/pkg/log"
	"moocss.com/gaea/pkg/trace"
)

// Recovery is a server middleware that recovers from any panics.
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Get(context.Background()).Println("recovery before")

		// 开始链路
		r, span := trace.StartSpan(r, "ServeHTTP")

		defer func() {
			if rec := recover(); rec != nil {
				ctx := r.Context()
				ctx = ctxkit.WithTraceID(ctx, trace.GetTraceID(ctx))
				log.Get(ctx).Error(rec, string(debug.Stack()))
			}
			span.Finish()
		}()

		next.ServeHTTP(w, r)

		log.Get(context.Background()).Println("recovery after")
	})
}
