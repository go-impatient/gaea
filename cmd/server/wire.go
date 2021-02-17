//+build wireinject

// The build tag makes sure the stub is not built in the final build.

package server

import (
	"github.com/google/wire"

	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/internal/data"
	"moocss.com/gaea/internal/service"
	"moocss.com/gaea/pkg/log"
)

//go:generate wire
func InitApp(logger log.Logger) (*service.Services, error) {
	panic(wire.Build(
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		wire.Struct(new(service.Services), "*"),
	))
}
