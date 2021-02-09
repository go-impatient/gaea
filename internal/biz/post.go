package biz

import (
	"context"

	"github.com/google/uuid"

	"moocss.com/gaea/internal/data/ent"
)

// UserRepository interface
type UserRepository interface {
	Exist(ctx context.Context, model *ent.User) (bool, error)
	List(ctx context.Context, limit, page int, sort string, model *ent.User) (total int, users []*ent.User, err error)
	Get(ctx context.Context, id uuid.UUID) (*ent.User, error)
	Create(ctx context.Context, model *ent.User) (*ent.User, error)
	Update(ctx context.Context, model *ent.User) (*ent.User, error)
	DeleteFull(ctx context.Context, model *ent.User) (*ent.User, error)
	Delete(ctx context.Context, id uuid.UUID) (*ent.User, error)
	Count(ctx context.Context) (int, error)
}

//NewUserUsecase
func NewUserUsecase(repo UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

//UserUsecase
type UserUsecase struct {
	repo UserRepository
}

func (uc *UserUsecase) CreateUser(ctx context.Context, u *ent.User) {
	uc.repo.Create(ctx, u)
}
