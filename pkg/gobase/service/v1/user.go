package v1

import (
	"context"
	v1 "github.com/Mr-LvGJ/gobase/pkg/gobase/model/v1"
	"github.com/Mr-LvGJ/gobase/pkg/gobase/store"
)

type userService struct {
	store store.Factory
}

func newUsers(srv *service) *userService {
	return &userService{store: srv.store}
}

func (u *userService) Create(ctx context.Context, user *v1.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Update(ctx context.Context, user *v1.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Delete(ctx context.Context, username string) error {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Get(ctx context.Context, username string) (*v1.User, error) {
	user, err := u.store.Users().Get(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) List(ctx context.Context) (*v1.UserList, error) {
	//TODO implement me
	panic("implement me")
}
