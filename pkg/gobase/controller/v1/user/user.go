package user

import (
	"github.com/Mr-LvGJ/gobase/pkg/common/core"
	"github.com/Mr-LvGJ/gobase/pkg/common/log"
	srvv1 "github.com/Mr-LvGJ/gobase/pkg/gobase/service/v1"
	"github.com/Mr-LvGJ/gobase/pkg/gobase/store"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	srv srvv1.Service
}

func NewUserController(store store.Factory) *UserController {
	return &UserController{
		srv: srvv1.NewService(store),
	}
}

func (u *UserController) Get(c *gin.Context) {
	log.Info("get user function called.")
	user, err := u.srv.Users().Get(c, c.Param("name"))
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, user)

}
