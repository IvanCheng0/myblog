package store

import (
	"context"

	"github.com/ivancheng/myblog/internal/pkg/model"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
}

type users struct {
	db *gorm.DB
}

// 确保 users 实现了 UserStore 接口.
var _ UserStore = (*users)(nil)

func newUser(db *gorm.DB) *users {
	return &users{db}
}

func (u *users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}
