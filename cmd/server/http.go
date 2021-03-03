package server

import (
	"net/http"

	"moocss.com/gaea/cmd/server/hooks"
	"moocss.com/gaea/internal/service"
	"moocss.com/gaea/pkg/log"

	blog_v1 "moocss.com/gaea/rpc/blog/v1"
	helloworld_v1 "moocss.com/gaea/rpc/helloworld/v1"
	user_v1 "moocss.com/gaea/rpc/user/v1"
)

func initMux(mux *http.ServeMux, services *service.Services, logger log.Logger, isInternal bool) {
	{
		server := services.HelloworldServer
		handler := helloworld_v1.NewHelloworldServer(server, hooks.Init(logger))
		mux.Handle(helloworld_v1.HelloworldPathPrefix, handler)
	}

	{
		server := services.UserServer
		handler := user_v1.NewUserServer(server, hooks.Init(logger))
		mux.Handle(user_v1.UserPathPrefix, handler)
	}

	{
		server := services.BlogServer
		handler := blog_v1.NewBlogServer(server, hooks.Init(logger))
		mux.Handle(blog_v1.BlogPathPrefix, handler)
	}
}

func initInternalMux(mux *http.ServeMux, services *service.Services, logger log.Logger) {
}
