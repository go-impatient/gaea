package recovery

import (
	"fmt"
	"net/http"
	"runtime"

	"moocss.com/gaea/pkg/ctxkit"
	"moocss.com/gaea/pkg/trace"
)

// Recovery is a server middleware that recovers from any panics.
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 开始链路
		r, span := trace.StartSpan(r, "ServeHTTP")

		defer func() {
			if rec := recover(); rec != nil {
				ctx := r.Context()
				ctx = ctxkit.WithTraceID(ctx, trace.GetTraceID(ctx))

				buf := make([]byte, 64<<10)
				n := runtime.Stack(buf, false)
				buf = buf[:n]

				fmt.Errorf("panic triggered: %v %s\n", rec, buf)
			}
			span.Finish()
		}()

		next.ServeHTTP(w, r)
	})
}
