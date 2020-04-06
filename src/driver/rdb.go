package driver

import (
	"fmt"

	app "github.com/sky0621/fs-mng-backend"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type RDB interface {
	// TODO: 抽象化検討
	GetDBWrapper() *sqlx.DB
	Close() error
}

func NewRDB(cfg app.Config) RDB {
	dsFormat := "dbname=%s user=%s password=%s sslmode=disable"
	dsn := fmt.Sprintf(dsFormat, cfg.RDBConfig.DBName, cfg.RDBConfig.User, cfg.RDBConfig.Password)
	dbWrapper, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err) // システム起動時なので
	}
	return &rdb{cfg: cfg, dbWrapper: dbWrapper}
}

type rdb struct {
	cfg       app.Config
	dbWrapper *sqlx.DB
}

func (r *rdb) GetDBWrapper() *sqlx.DB {
	return r.dbWrapper
}

func (r *rdb) Close() error {
	if r == nil {
		return nil
	}
	if r.dbWrapper == nil {
		return nil
	}
	return r.dbWrapper.Close()
}
