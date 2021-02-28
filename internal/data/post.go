package data

import (
	"context"

	"github.com/pkg/errors"

	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/internal/data/ent"
	"moocss.com/gaea/pkg/log"
)

type postRepo struct {
	data *Data
	log  *log.Helper
}

// NewPostRepo returns an instance of `postRepo`.
func NewPostRepo(data *Data, logger log.Logger) biz.PostRepo {
	return &postRepo{
		data: data,
		log:  log.NewHelper("post_repo", logger),
	}
}

// implement biz.PostRepo
func (r *postRepo) Create(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	r.log.Info("Received PostRepository.Create")
	out, err := r.data.db.Post.Create().
		SetTitle(post.Title).
		SetContent(post.Content).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	do := biz.PostToDO(out)

	return do, nil
}

func (r *postRepo) Update(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	r.log.Info("Received PostRepository.Update")

	var tx *ent.Tx
	tx, err := r.data.db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	po := biz.PostToPO(post)

	updateOp := tx.Post.UpdateOne(po)

	out, err := updateOp.Save(ctx)
	do := biz.PostToDO(out)
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
