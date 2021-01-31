package server

import (
	"net/http"

	"moocss.com/gaea/cmd/server/hook"
	"moocss.com/gaea/pkg/twirp"

	example_v1 "moocss.com/gaea/rpc/example/v1"
)

var hooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewLog(),
)

func initMux(mux *http.ServeMux, isInternal bool) {

	{
		server := &example_v1.HelloworldServer{}
		handler := example_v1.NewHelloworldServer(server, hooks)
		mux.Handle(example_v1.HelloworldPathPrefix, handler)
	}
}

func initInternalMux(mux *http.ServeMux) {
}
