package v1

import (
	"context"
	"regexp"
	"sync"

	"github.com/Mr-LvGJ/gobase/pkg/common/auth"
	"github.com/Mr-LvGJ/gobase/pkg/common/errno"
	"github.com/Mr-LvGJ/gobase/pkg/common/log"
	metav1 "github.com/Mr-LvGJ/gobase/pkg/gobase/meta/v1"

	"github.com/Mr-LvGJ/gobase/pkg/common/auth"
	"github.com/Mr-LvGJ/gobase/pkg/common/errno"
	"github.com/Mr-LvGJ/gobase/pkg/common/log"
	metav1 "github.com/Mr-LvGJ/gobase/pkg/gobase/meta/v1"
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
	if err := u.store.Users().Create(ctx, user); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.* for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}
	return nil
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

func (u *userService) List(ctx context.Context, options metav1.ListOptions) (*v1.UserList, error) {
	users, err := u.store.Users().List(ctx, options)
	if err != nil {
		log.Error("list users from storage failed", err.Error())
		return nil, err
	}

	wg := sync.WaitGroup{}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)
	var m sync.Map

	for _, user := range users.Items {
		wg.Add(1)
		go func(user *v1.User) {
			defer wg.Done()

			shadowedPassword, err := auth.Shadow(user.Password)
			if err != nil {
				errChan <- err
				return
			}
			user.Password = shadowedPassword
			m.Store(user.ID, user)
		}(user)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}
	infos := make([]*v1.User, 0, len(users.Items))
	for _, user := range users.Items {
		info, _ := m.Load(user.ID)
		infos = append(infos, info.(*v1.User))
	}
	log.Info("get info from backend.")
	return &v1.UserList{Items: infos}, nil
}
