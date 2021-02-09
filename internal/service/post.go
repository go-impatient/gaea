package service

import (
	"context"

	"moocss.com/gaea/internal/biz"
	pb "moocss.com/gaea/rpc/blog/v1"
)

type PostServer struct {
	puc *biz.PostUsecase
}

func NewPostServer(pr biz.PostRepository) *PostServer {
	return &PostServer{
		puc: biz.NewPostUsecase(pr),
	}
}

func (s *PostServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	// dto -> do
	post := &biz.Post{
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}

	s.puc.CreatePost(ctx, post)

	return &pb.CreatePostReply{}, nil
}

func (s *PostServer) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostReply, error) {
	panic("implement me")
}

func (s *PostServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostReply, error) {
	panic("implement me")
}

func (s *PostServer) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostReply, error) {
	panic("implement me")
}

func (s *PostServer) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostReply, error) {
	panic("implement me")
}
