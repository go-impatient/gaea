package biz

import (
	"context"
	"time"
)

// User .
type User struct {
	ID                    uint64
	Username              string    // 用户名
	Password              string    // 密码
	Salt                  string    // 盐
	Realname              string    // 真实姓名
	Nickname              string    // 昵称
	Email                 string    // 邮箱
	Phone                 string    // 手机号码
	Sex                   int       // 性别: 0 男性, 1 女性, 2未知
	Status                int       // 状态: 1启用, 0禁用
	IsDeleted             int       // 是否已删除 : 1删除, 0未删除
	CreateUser            string    // 创建人
	UpdateUser            string    // 修改人
	PasswordErrorLastTime time.Time // 最后一次输错密码时间
	PasswordErrorNum      int       // 密码错误次数
	PasswordExpireTime    time.Time // 密码过期时间
}

// UserRepo interface
type UserRepo interface {
	Exist(ctx context.Context, user *User) (bool, error)
	List(ctx context.Context, limit, page int, sort string, user *User) (total int, users []*User, err error)
	Get(ctx context.Context, id int) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	DeleteFull(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id int) (*User, error)
	Count(ctx context.Context) (int, error)
}

// UserUsecase .
type UserUsecase struct {
	repo UserRepo
}

// NewUserUsecase .
func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// Create .
func (uc *UserUsecase) Create(ctx context.Context, user *User) (*User, error) {
	return uc.repo.Create(ctx, user)
}

// Update .
func (uc *UserUsecase) Update(ctx context.Context, user *User) (*User, error) {
	return uc.repo.Update(ctx, user)
}
