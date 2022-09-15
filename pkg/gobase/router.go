package gobase

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/Mr-LvGJ/gobase/pkg/common/core"
	"github.com/Mr-LvGJ/gobase/pkg/common/errno"
	"github.com/Mr-LvGJ/gobase/pkg/common/middleware"
	"github.com/Mr-LvGJ/gobase/pkg/gobase/controller/v1/user"
	"github.com/Mr-LvGJ/gobase/pkg/gobase/store"
)

func LoadRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine, mw ...gin.HandlerFunc) {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middleware.RequestID())
}

func installController(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.InternalServerError, nil)
	})

	g.GET("/healthz", func(c *gin.Context) {
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	dbClient := store.Client()
	userController := user.NewUserController(dbClient)

	pprof.Register(g)

	g.POST("/login", userController.Login)

	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userv1.GET(":name", userController.Get)
		}
	}

}
