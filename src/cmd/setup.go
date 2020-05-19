package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/boil"
)

type closeDBFunc func() error

func newDB(e *env) (*sql.DB, closeDBFunc) {
	var db *sql.DB
	closeDBFunc := func() error {
		if db == nil {
			return nil
		}
		if err := db.Close(); err != nil {
			return err
		}
		return nil
	}
	{
		var err error
		// MEMO: ひとまずローカルのコンテナ相手の接続前提なので、べたに書いておく。
		dataSourceName := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s port=%s", e.DBName, e.DBUser, e.DBPassword, e.DBSSLMode, e.DBPort)
		db, err = sql.Open(e.DBDriverName, dataSourceName)
		if err != nil {
			panic(err)
		}

		boil.DebugMode = true

		var loc *time.Location
		loc, err = time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic(err)
		}
		boil.SetLocation(loc)
	}
	return db, closeDBFunc
}
