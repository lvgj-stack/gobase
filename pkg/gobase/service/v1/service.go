package v1

import "github.com/Mr-LvGJ/gobase/pkg/gobase/store"

type service struct {
	store store.Factory
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}

func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}
