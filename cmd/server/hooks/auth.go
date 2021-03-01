package hooks

import (
	"context"
	"fmt"

	"moocss.com/gaea/pkg/twirp"
)

// Auth 用户认证
func NewAuth() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			// 获取当前用户的token
			token := ctx.Value("GAEA_TOKEN").(string)
			fmt.Println("token:", token)

			return ctx, nil
		},
	}
}
