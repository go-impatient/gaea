package blog_v1

import (
	"context"
)

// PostServer 实现 /blog.v1.Post 服务
// FIXME 服务必须写注释
type PostServer struct{}

func (s *PostServer) CreatePost(ctx context.Context, request *CreatePostRequest) (*CreatePostReply, error) {
	panic("implement me")
}

func (s *PostServer) UpdatePost(ctx context.Context, request *UpdatePostRequest) (*UpdatePostReply, error) {
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
