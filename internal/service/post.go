package service

import (
	"context"

	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/rpc/blog/v1"
)

type PostServer struct {
	u *biz.UserUsecase
}

func NewUserService(u *biz.UserUsecase)  *PostServer  {
	return &PostServer{
		u: u,
	}
}

func (s *PostServer) CreatePost(ctx context.Context, request *blog_v1.CreatePostRequest) (*blog_v1.CreatePostReply, error) {
	panic("implement me")
}

func (s *PostServer) UpdatePost(ctx context.Context, request *blog_v1.UpdatePostRequest) (*blog_v1.UpdatePostReply, error) {
	panic("implement me")
}

func (s *PostServer) DeletePost(ctx context.Context, request *DeletePostRequest) (*DeletePostReply, error) {
	panic("implement me")
}

func (s *PostServer) GetPost(ctx context.Context, request *GetPostRequest) (*GetPostReply, error) {
	panic("implement me")
}

func (s *PostServer) ListPost(ctx context.Context, request *ListPostRequest) (*ListPostReply, error) {
	panic("implement me")
}
