package data

import (
	"context"
	"time"

	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/pkg/log"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo returns an instance of `articleRepo`.
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper("article_repo", logger),
	}
}

func (r *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	ps, err := r.data.db.Article.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Article, 0)

	for _, p := range ps {
		rv = append(rv, biz.ArticleToDO(p))
	}
	return rv, nil
}

func (r *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	p, err := r.data.db.Article.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	do := biz.ArticleToDO(p)
	return do, nil
}

func (r *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	_, err := r.data.db.Article.
		Create().
		SetTitle(article.Title).
		SetContent(article.Content).
		Save(ctx)
	return err
}

func (r *articleRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) error {
	p, err := r.data.db.Article.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetTitle(article.Title).
		SetContent(article.Content).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	return err
}

func (r *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	return r.data.db.Article.DeleteOneID(id).Exec(ctx)
}

/*
func (r *articleRepo) UpdateArticle(ctx context.Context, post *biz.Article) (*biz.Article, error) {
	r.log.Info("Received ArticleRepository.Update")

	var tx *ent.Tx
	tx, err := r.data.db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	po := biz.ArticleToPO(post)

	updateOp := tx.Article.UpdateOne(po)

	out, err := updateOp.Save(ctx)
	do := biz.ArticleToDO(out)
	if err != nil {
		err = rollback(tx, err)
		return nil, err
	}

	err = tx.Commit()

	return do, nil
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
	}
	return err
}
*/
