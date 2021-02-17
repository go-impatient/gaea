package service

import (
	"context"
	pb "moocss.com/gaea/rpc/helloworld/v1"
)

// HelloworldServer 实现 /helloworld.v1.Helloworld 服务
// FIXME 服务必须写注释
type HelloworldServer struct{}

func NewHelloworldServer() *HelloworldServer {
	return &HelloworldServer{}
}

// Echo 实现 /helloworld.v1.Helloworld/Echo 接口
// FIXME 接口必须写注释
//
// 这里的行尾注释 sniper:foo 有特殊含义，是可选的
// 框架会将此处冒号后面的值(foo)注入到 ctx 中，
// 用户可以使用 twirp.MethodOption(ctx) 查询，并执行不同的逻辑
// 这个 sniper 前缀可以通过 --twirp_out=option_prefix=sniper:. 自定义
func (s *HelloworldServer) Echo(ctx context.Context, req *pb.HelloworldEchoReq) (resp *pb.HelloworldEchoResp, err error) {
	return &pb.HelloworldEchoResp{Msg: req.Msg}, nil
}
