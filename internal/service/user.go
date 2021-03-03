package service

import (
	"context"

	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/pkg/log"
	pb "moocss.com/gaea/rpc/user/v1"
)

// NewUserServer .
func NewUserServer(user *biz.UserUsecase, logger log.Logger) *UserServer {
	return &UserServer{
		user: user,
		log:  log.NewHelper("post", logger),
	}
}

// Echo 实现 /user.v1.User/Echo 接口
// FIXME 接口必须写注释
//
// 这里的行尾注释 sniper:foo 有特殊含义，是可选的
// 框架会将此处冒号后面的值(foo)注入到 ctx 中，
// 用户可以使用 twirp.MethodOption(ctx) 查询，并执行不同的逻辑
// 这个 sniper 前缀可以通过 --twirp_out=option_prefix=sniper:. 自定义
func (s *UserServer) Echo(ctx context.Context, req *pb.UserEchoReq) (resp *pb.UserEchoResp, err error) {
	return &pb.UserEchoResp{Msg: req.Msg}, nil
}
