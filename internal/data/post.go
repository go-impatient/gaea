package data

import (
	"context"

	"github.com/google/uuid"

	"moocss.com/gaea/internal/biz"
	"moocss.com/gaea/internal/data/ent"
)

var _ biz.UserRepository = (biz.UserRepository)(nil)

// userRepository struct
type userRepository struct {
	dbClient *ent.Client
}

// NewUserRepository returns an instance of `UserRepository`.
func NewUserRepository(dbClient *ent.Client) biz.UserRepository {
	return &userRepository{
		dbClient: dbClient,
	}
}

func (u *userRepository) Exist(ctx context.Context, model *ent.User) (bool, error) {
	panic("implement me")
}

func (u *userRepository) List(ctx context.Context, limit, page int, sort string, model *ent.User) (total int, users []*ent.User, err error) {
	panic("implement me")
}

func (u *userRepository) Get(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	panic("implement me")
}

func (u *userRepository) Create(ctx context.Context, model *ent.User) (*ent.User, error) {
	panic("implement me")
}

func (u *userRepository) Update(ctx context.Context, model *ent.User) (*ent.User, error) {
	panic("implement me")
}

func (u *userRepository) DeleteFull(ctx context.Context, model *ent.User) (*ent.User, error) {
	panic("implement me")
}

func (u userRepository) Delete(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	panic("implement me")
}

func (u *userRepository) Count(ctx context.Context) (int, error) {
	panic("implement me")
}
