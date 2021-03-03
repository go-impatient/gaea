package biz

import (
	"context"
	"time"

	"moocss.com/gaea/internal/data/ent"
)

// Article .
type Article struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Like      int64
}

// ArticleToPO article DO -> PO
func ArticleToPO(article *Article) *ent.Article {
	return &ent.Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		DeletedAt: article.DeletedAt,
	}
}

// ArticleToDO article PO -> DO
func ArticleToDO(article *ent.Article) *Article {
	return &Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		DeletedAt: article.DeletedAt,
	}
}

// ArticleRepo interface
type ArticleRepo interface {
	// db
	ListArticle(ctx context.Context) ([]*Article, error)
	GetArticle(ctx context.Context, id int64) (*Article, error)
	CreateArticle(ctx context.Context, article *Article) error
	UpdateArticle(ctx context.Context, id int64, article *Article) error
	DeleteArticle(ctx context.Context, id int64) error

	// redis
	GetArticleLike(ctx context.Context, id int64) (rv int64, err error)
	IncArticleLike(ctx context.Context, id int64) error
}

// ArticleUsecase .
type ArticleUsecase struct {
	repo ArticleRepo
}

// NewArticleUsecase .
func NewArticleUsecase(repo ArticleRepo) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

// List .
func (uc *ArticleUsecase) List(ctx context.Context) (ps []*Article, err error) {
	ps, err = uc.repo.ListArticle(ctx)
	if err != nil {
		return
	}
	return
}

// Get .
func (uc *ArticleUsecase) Get(ctx context.Context, id int64) (p *Article, err error) {
	p, err = uc.repo.GetArticle(ctx, id)
	if err != nil {
		return
	}
	err = uc.repo.IncArticleLike(ctx, id)
	if err != nil {
		return
	}
	p.Like, err = uc.repo.GetArticleLike(ctx, id)
	if err != nil {
		return
	}
	return
}

// Create .
func (uc *ArticleUsecase) Create(ctx context.Context, article *Article) error {
	return uc.repo.CreateArticle(ctx, article)
}

// Update .
func (uc *ArticleUsecase) Update(ctx context.Context, id int64, article *Article) error {
	return uc.repo.UpdateArticle(ctx, id, article)
}

// Delete .
func (uc *ArticleUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteArticle(ctx, id)
}
