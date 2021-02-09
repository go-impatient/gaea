package data

import (
	"context"

	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/internal/data/ent"
)

var _ biz.PostRepository = (biz.PostRepository)(nil)

// postRepository struct
type postRepository struct {
	db *ent.PostClient
}

// NewPostRepository returns an instance of `UserRepository`.
func NewPostRepository(c *ent.Client) biz.PostRepository {
	return &postRepository{
		db: c.Post,
	}
}

// implement biz.PostRepository
func (p *postRepository) List(ctx context.Context, limit, page int, sort string, model *biz.Post) (total int, users []*biz.Post, err error) {
	return 0, nil, nil
}

func (p *postRepository) Get(ctx context.Context, id int) (*biz.Post, error) {
	return nil, nil
}

func (p *postRepository) Create(ctx context.Context, model *biz.Post) (*biz.Post, error) {
	return nil, nil
}

func (p *postRepository) Update(ctx context.Context, model *biz.Post) (*biz.Post, error) {
	return nil, nil
}

func (p postRepository) DeleteFull(ctx context.Context, model *biz.Post) (*biz.Post, error) {
	return nil, nil
}

func (p postRepository) Delete(ctx context.Context, id int) (*biz.Post, error) {
	return nil, nil
}

func (p postRepository) Count(ctx context.Context) (int, error) {
	return 0, nil
}
