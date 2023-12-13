package store

import (
	"context"

	"gorm.io/gorm"

	"github.com/Mr-LvGJ/jota/log"
	"github.com/Mr-LvGJ/jota/models"

	"github.com/Mr-LvGJ/gobase/pkg/common/setting"
)

var (
	client Factory
)

func Client() Factory {
	return client
}

type datastore struct {
	db *gorm.DB
}

func (d *datastore) Close() error {
	//TODO implement me
	panic("implement me")
}

func (d *datastore) Users() UserStore {
	return newUsers(d)
}

func Setup(ctx context.Context) (*gorm.DB, error) {
	log.Info(ctx, "init database", "config", setting.C().Database)
	db, err := models.New(setting.C().Database)
	if err != nil {
		panic(err)
	}
	client = &datastore{
		db: db,
	}
	return db, nil
}
