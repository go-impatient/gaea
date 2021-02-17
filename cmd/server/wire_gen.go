// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"github.com/sirupsen/logrus"
	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/internal/data"
	"moocss.com/gaea/internal/service"
)

// Injectors from wire.go:

func InitApp(logger2 *logrus.Entry) (*service.Services, error) {
	helloworldServer := service.NewHelloworldServer()
	dataData, err := data.NewData(logger2)
	if err != nil {
		return nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger2)
	userUsecase := biz.NewUserUsecase(userRepo)
	userServer := service.NewUserServer(userUsecase, logger2)
	postRepo := data.NewPostRepo(dataData, logger2)
	postUsecase := biz.NewPostUsecase(postRepo)
	postServer := service.NewPostServer(postUsecase, logger2)
	services := &service.Services{
		HelloworldServer: helloworldServer,
		UserServer:       userServer,
		PostServer:       postServer,
	}
	return services, nil
}
