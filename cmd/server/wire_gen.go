// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/internal/data"
	"moocss.com/gaea/internal/service"
	"moocss.com/gaea/pkg/log"
)

// Injectors from wire.go:

func InitApp(logger log.Logger) (*service.Services, error) {
	helloworldServer := service.NewHelloworldServer()
	dataData, err := data.NewData(logger)
	if err != nil {
		return nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo)
	userServer := service.NewUserServer(userUsecase, logger)
	postRepo := data.NewPostRepo(dataData, logger)
	postUsecase := biz.NewPostUsecase(postRepo)
	postServer := service.NewPostServer(postUsecase, logger)
	services := &service.Services{
		HelloworldServer: helloworldServer,
		UserServer:       userServer,
		PostServer:       postServer,
	}
	return services, nil
}
