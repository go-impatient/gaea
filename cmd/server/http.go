package server

import (
	"net/http"

	"moocss.com/gaea/cmd/server/hook"
	"moocss.com/gaea/pkg/twirp"

	blog_v1 "moocss.com/gaea/rpc/blog/v1"
	example_v1 "moocss.com/gaea/rpc/example/v1"
)

var hooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewLog(),
)

// var authHooks = twirp.ChainHooks(
// 	hook.NewHeaders(),
// 	hook.NewRequestID(),
// 	hook.NewLog(),
// 	hook.NewAuth(),
// )

func initMux(mux *http.ServeMux, isInternal bool) {

	{
		server := &example_v1.HelloworldServer{}
		handler := example_v1.NewHelloworldServer(server, hooks)
		mux.Handle(example_v1.HelloworldPathPrefix, handler)
	}

	{
		server := &blog_v1.PostServer{}
		handler := blog_v1.NewPostServer(server, hooks)
		mux.Handle(blog_v1.PostPathPrefix, handler)
	}
}

func initInternalMux(mux *http.ServeMux) {
}
