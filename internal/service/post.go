package service

import (
	"context"

	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/pkg/log"
	pb "moocss.com/gaea/rpc/blog/v1"
)

type PostServer struct {
	post *biz.PostUsecase

	log *log.Helper
}

func NewPostServer(post *biz.PostUsecase, logger log.Logger) *PostServer {
	return &PostServer{
		post: post,
		log:  log.NewHelper("post", logger),
	}
}

func (s *PostServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	// dto -> do
	post := &biz.Post{
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}

	s.post.CreatePost(ctx, post)

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
