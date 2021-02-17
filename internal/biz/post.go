package biz

import (
	"context"
	"time"

	"moocss.com/gaea/internal/data/ent"
)

type Post struct {
	ID      uint64
	Title   string
	Content string
	CTime   time.Time
}

// post DO -> PO
func PostToPO(post *Post) *ent.Post {
	return &ent.Post{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	}
}

// post PO -> DO
func PostToDO(post *ent.Post) *Post {
	return &Post{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		CTime:   post.Ctime,
	}
}

// PostRepo interface
type PostRepo interface {
	Create(ctx context.Context, model *Post) (*Post, error)
	Update(ctx context.Context, model *Post) (*Post, error)
}

// PostUsecase
type PostUsecase struct {
	repo PostRepo
}

// NewPostUsecase
func NewPostUsecase(repo PostRepo) *PostUsecase {
	return &PostUsecase{repo: repo}
}

func (uc *PostUsecase) CreatePost(ctx context.Context, post *Post) (*Post, error) {
	return uc.repo.Create(ctx, post)
}

func (uc *PostUsecase) UpdatePost(ctx context.Context, post *Post) (*Post, error) {
	return uc.repo.Update(ctx, post)
}
