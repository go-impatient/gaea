package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserServer, NewPostServer)

type Services struct {
	UserServer *UserServer
	PostServer *PostServer
}

