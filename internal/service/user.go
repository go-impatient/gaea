package service

import (
	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/pkg/log"
)

type UserServer struct {
	user *biz.UserUsecase

	log log.Logger
}

func NewUserServer(user *biz.UserUsecase, logger log.Logger) *UserServer {
	return &UserServer{
		user: user,
		log:  logger,
	}
}
