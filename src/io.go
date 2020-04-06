package system

import (
	"fmt"

	multierror "github.com/hashicorp/go-multierror"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type IO interface {
	Close() error
}

func NewIO(rdb RDB) IO {
	return &io{rdb: rdb}
}

type io struct {
	rdb RDB
}

func (i *io) Close() error {
	if i == nil {
		return nil
	}
	var err error
	if i.rdb != nil {
		err = multierror.Append(err, i.rdb.Close())
	}
	return err
}

type RDB interface {
	Close() error
}

func NewRDB(cfg Config) RDB {
	dsFormat := "dbname=%s user=%s password=%s sslmode=disable"
	dsn := fmt.Sprintf(dsFormat, cfg.DBName, cfg.DBUser, cfg.DBPassword)
	dbWrapper, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err) // システム起動時なので
	}
	return &rdb{dbWrapper: dbWrapper}
}

type rdb struct {
	dbWrapper *sqlx.DB
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
