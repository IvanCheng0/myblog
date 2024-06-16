package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/ivancheng/myblog/internal/myblog/biz"
	"github.com/ivancheng/myblog/internal/myblog/store"
	"github.com/ivancheng/myblog/internal/pkg/core"
	"github.com/ivancheng/myblog/internal/pkg/errno"
	"github.com/ivancheng/myblog/internal/pkg/log"
	v1 "github.com/ivancheng/myblog/pkg/api/myblog/v1"
)

type UserController struct {
	b biz.IBiz
}

func New(ds store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}

func (u *UserController) Create(c *gin.Context) {
	log.C(c).Infow("Create user", "req", c.Request.Body)

	var req v1.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(&req); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := u.b.Users().Create(c, &req); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}
