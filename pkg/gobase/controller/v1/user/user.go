package user

import (
	"github.com/Mr-LvGJ/gobase/pkg/common/auth"
	"github.com/Mr-LvGJ/gobase/pkg/common/errno"
	"github.com/Mr-LvGJ/gobase/pkg/common/token"
	"github.com/gin-gonic/gin"

	"github.com/Mr-LvGJ/gobase/pkg/common/core"
	"github.com/Mr-LvGJ/gobase/pkg/common/log"
	srvv1 "github.com/Mr-LvGJ/gobase/pkg/gobase/service/v1"
	"github.com/Mr-LvGJ/gobase/pkg/gobase/store"
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

func (u *UserController) Login(c *gin.Context) {
	var r LoginRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}
	log.Info("LoginRequest", "request", r, "username", r.Username)

	user, err := u.srv.Users().Get(c, r.Username)
	if err != nil {
		core.WriteResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err = auth.Compare(user.Password, r.Password); err != nil {
		core.WriteResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(r.Username)
	if err != nil {
		core.WriteResponse(c, errno.ErrToken, nil)
		return
	}

	core.WriteResponse(c, nil, LoginResponse{
		t,
	})

}
