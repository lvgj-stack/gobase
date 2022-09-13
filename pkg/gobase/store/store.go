package store

import (
	"fmt"
	"sync"
	"time"

	"github.com/Mr-LvGJ/gobase/pkg/common/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	client Factory
	once   sync.Once
)

func Client() Factory {
	return client

}

func SetClient(factory Factory) {
	client = factory

}

type Options struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
	Logger                logger.Interface
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

func New(opts *Options) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Database,
		true,
		"Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: opts.Logger,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
	return db, nil

}

func Setup() (Factory, error) {
	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		options := &Options{
			Host:                  setting.C().Database.Host,
			Username:              setting.C().Database.Username,
			Password:              setting.C().Database.Password,
			Database:              setting.C().Database.DatabaseName,
			MaxIdleConnections:    setting.C().Database.MaxIdleConns,
			MaxOpenConnections:    setting.C().Database.MaxOpenConns,
			MaxConnectionLifeTime: setting.C().Database.ConnMaxLifetime,
			LogLevel:              setting.C().Database.LoggerLevel,
		}
		dbIns, err = New(options)
		client = &datastore{
			db: dbIns,
		}
	})
	if client == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", client, err)
	}
	return client, nil
}
