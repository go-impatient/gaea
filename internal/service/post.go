package service

import (
	"context"

	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/rpc/blog/v1"
)

type PostService struct {
	puc *biz.PostUsecase
}

func NewPostService(pr biz.PostRepository) *PostService {
	return &PostService{
		puc: biz.NewPostUsecase(pr),
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *blog_v1.CreatePostRequest) (*blog_v1.CreatePostReply, error) {
	// dto -> do
	post := &biz.Post {
		Title: req.GetTitle(),
		Content: req.GetContent(),
	}

	s.puc.CreatePost(ctx, post)

	return &blog_v1.CreatePostReply{}, nil
}

func (s *PostService) UpdatePost(ctx context.Context, req *blog_v1.UpdatePostRequest) (*blog_v1.UpdatePostReply, error) {
	panic("implement me")
}

func (s *PostService) DeletePost(ctx context.Context, req *blog_v1.DeletePostRequest) (*blog_v1.DeletePostReply, error) {
	panic("implement me")
}

func (s *PostService) GetPost(ctx context.Context, req *blog_v1.GetPostRequest) (*blog_v1.GetPostReply, error) {
	panic("implement me")
}

func (s *PostService) ListPost(ctx context.Context, req *blog_v1.ListPostRequest) (*blog_v1.ListPostReply, error) {
	panic("implement me")
}
