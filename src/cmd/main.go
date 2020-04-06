package main

import (
	"os"
	"os/signal"
	"syscall"

	system "github.com/sky0621/fs-mng-backend"
)

type exitCode int

const (
	normalEnd   = 0
	abnormalEnd = -1
)

func main() {
	os.Exit(int(execMain()))
}

func execMain() exitCode {
	cfg := system.Config{
		RDBConfig: system.RDBConfig{
			DBName:   os.Getenv("FIKTIVT_RDB_DBNAME"),
			User:     os.Getenv("FIKTIVT_RDB_USER"),
			Password: os.Getenv("FIKTIVT_RDB_PASSWORD"),
		},
		WebConfig: system.WebConfig{ListenPort: os.Getenv("FIKTIVT_WEB_LISTENPORT")},
	}

	app := di(cfg)
	defer app.Shutdown()

	// OSシグナル受信したらグレースフルシャットダウン
	go func() {
		q := make(chan os.Signal)
		signal.Notify(q, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-q

		// TODO: アプリ終了前の後始末を実装！

		os.Exit(int(abnormalEnd))
	}()

	return normalEnd
}
