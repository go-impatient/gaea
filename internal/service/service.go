package service

import (
	"github.com/google/wire"

	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/pkg/log"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewHelloworldServer, NewUserServer, NewBlogServer)

// UserServer 实现 /user.v1.User 服务
type UserServer struct {
	log *log.Helper

	user *biz.UserUsecase
}

// BlogServer 实现 /blog.v1.Blog 服务
type BlogServer struct {
	log *log.Helper

	article *biz.ArticleUsecase
}

// Services .
type Services struct {
	HelloworldServer *HelloworldServer
	UserServer       *UserServer
	BlogServer       *BlogServer
}
