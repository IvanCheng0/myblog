package user

// package biz

import (
	"context"
	"regexp"

	// "github.com/gin-gonic/gin"
	"github.com/ivancheng/myblog/internal/myblog/store"
	"github.com/ivancheng/myblog/internal/pkg/errno"
	"github.com/ivancheng/myblog/internal/pkg/model"
	v1 "github.com/ivancheng/myblog/pkg/api/myblog/v1"
	"github.com/jinzhu/copier"
)

// UserBiz 定义了 user 模块在 biz 层所实现的方法.
type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
}

type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)

func New(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}

func (b *userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, r)

	if err := b.ds.Users().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}

	return nil
}
