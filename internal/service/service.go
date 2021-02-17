package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewHelloworldServer, NewUserServer, NewPostServer)

type Services struct {
	HelloworldServer *HelloworldServer
	UserServer *UserServer
	PostServer *PostServer
}

