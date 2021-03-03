package service

import (
	"context"

	"go.opentelemetry.io/otel"

	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/pkg/log"
	pb "moocss.com/gaea/rpc/blog/v1"
)

func NewBlogServer(article *biz.ArticleUsecase, logger log.Logger) *BlogServer {
	return &BlogServer{
		article: article,
		log:     log.NewHelper("article", logger),
	}
}

func (s *BlogServer) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	s.log.Infof("input data %v", req)
	err := s.article.Create(ctx, &biz.Article{
		Title:   req.Title,
		Content: req.Content,
	})
	return &pb.CreateArticleReply{}, err
}

func (s *BlogServer) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	s.log.Infof("input data %v", req)
	err := s.article.Update(ctx, req.Id, &biz.Article{
		Title:   req.Title,
		Content: req.Content,
	})
	return &pb.UpdateArticleReply{}, err
}

func (s *BlogServer) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	s.log.Infof("input data %v", req)
	err := s.article.Delete(ctx, req.Id)
	return &pb.DeleteArticleReply{}, err
}

func (s *BlogServer) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	tr := otel.Tracer("api")
	_, span := tr.Start(ctx, "GetArticle")
	defer span.End()
	p, err := s.article.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetArticleReply{Article: &pb.Article{Id: p.ID, Title: p.Title, Content: p.Content, Like: p.Like}}, nil
}

func (s *BlogServer) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	ps, err := s.article.List(ctx)
	reply := &pb.ListArticleReply{}
	for _, p := range ps {
		reply.Results = append(reply.Results, &pb.Article{
			Id:      p.ID,
			Title:   p.Title,
			Content: p.Content,
		})
	}
	return reply, err
}
