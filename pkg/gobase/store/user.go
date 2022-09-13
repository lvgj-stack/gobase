package store

import (
	"context"
	"errors"
	v1 "github.com/Mr-LvGJ/gobase/pkg/gobase/model/v1"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{ds.db}

}

func (u *users) Create(ctx context.Context, user *v1.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *users) Update(ctx context.Context, user *v1.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *users) Delete(ctx context.Context, username string) error {
	//TODO implement me
	panic("implement me")
}

func (u *users) Get(ctx context.Context, username string) (*v1.User, error) {
	user := &v1.User{}
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return user, nil
}

func (u *users) List(ctx context.Context) (*v1.UserList, error) {
	//TODO implement me
	panic("implement me")
}
