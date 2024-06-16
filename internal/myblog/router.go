package myblog

// package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ivancheng/myblog/internal/myblog/controller/v1/user"
	"github.com/ivancheng/myblog/internal/myblog/store"
	"github.com/ivancheng/myblog/internal/pkg/core"
	"github.com/ivancheng/myblog/internal/pkg/errno"
	"github.com/ivancheng/myblog/internal/pkg/log"
)

func installRouters(g *gin.Engine) error {
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	uc := user.New(store.S)

	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
		}

	}
	return nil
}
