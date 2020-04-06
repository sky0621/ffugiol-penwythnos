//+build wireinject

package main

import (
	system "github.com/sky0621/fs-mng-backend"
	"github.com/sky0621/fs-mng-backend/driver"
	"github.com/sky0621/fs-mng-backend/usecase"

	"github.com/google/wire"
)

var superSet = wire.NewSet(
	// ロガー
	system.NewLogger,

	// Config => RDBコネクションプール
	driver.NewRDB,

	// RDBコネクションプール => domain層インタフェースをadapter層で実装
	//gateway.NewItem,
	//gateway.NewItemHolder,

	// domain層インタフェース => usecase層
	usecase.NewItem,
	usecase.NewItemHolder,

	// usecase層 => GraphQLリゾルバー
	//controller.NewResolverRoot,

	// WebFramework
	driver.NewWeb,

	// システムそのもの
	system.NewApp,
)

func di(cfg system.Config) App {
	wire.Build(superSet)
	return &AppImpl{}
}
