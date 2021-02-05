package hook

import (
	"context"

	"moocss.com/gaea/pkg/twirp"
)

// NewHeaders
func NewHeaders() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			request, _ := twirp.HttpRequest(ctx)

			token := request.Header.Get("GAEA_TOKEN")

			// 注入用户token
			ctx = context.WithValue(ctx, "token", token)

			err := twirp.SetHTTPResponseHeader(ctx, "Token-Lifecycle", "60")
			if err != nil {
				return nil, twirp.InternalErrorWith(err)
			}

			// 注入 用户设备  iso、android、web
			ctx = context.WithValue(ctx, "token", request.Header.Get("GAEA_DEVICE"))
			// 注入 版本标识
			ctx = context.WithValue(ctx, "token", request.Header.Get("GAEA_VERSION"))

			return ctx, nil
		},
	}
}
