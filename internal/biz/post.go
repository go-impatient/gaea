package biz

import (
	"context"
	"time"
)

type Post struct {
	ID int64
	Title string
	Content string
	CTime time.Time
}

// PostRepository interface
type PostRepository interface {
	List(ctx context.Context, limit, page int, sort string, model *Post) (total int, users []*Post, err error)
	Get(ctx context.Context, id int) (*Post, error)
	Create(ctx context.Context, model *Post) (*Post, error)
	Update(ctx context.Context, model *Post) (*Post, error)
	DeleteFull(ctx context.Context, model *Post) (*Post, error)
	Delete(ctx context.Context, id int) (*Post, error)
	Count(ctx context.Context) (int, error)
}

// NewPostUsecase
func NewPostUsecase(repo PostRepository) *PostUsecase {
	return &PostUsecase{repo: repo}
}

//PostUsecase
type PostUsecase struct {
	repo PostRepository
}

func (uc *PostUsecase) CreatePost(ctx context.Context, post *Post) {
	uc.repo.Create(ctx, post)
}
