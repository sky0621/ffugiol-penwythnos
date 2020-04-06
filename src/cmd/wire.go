//+build wireinject

package main

import (
	system "github.com/sky0621/fs-mng-backend"
	"github.com/sky0621/fs-mng-backend/usecase"

	"github.com/google/wire"
)

var superSet = wire.NewSet(
	// config => RDBコネクションプール等
	system.NewRDB,
	system.NewIO,

	// RDBコネクションプール => domain層インタフェースをadapter層で実装
	//gateway.NewItem,
	//gateway.NewItemHolder,

	// domain層インタフェース+IO => usecase層
	usecase.NewItem,
	usecase.NewItemHolder,

	// usecase層 => GraphQLリゾルバー
	//controller.NewResolverRoot,

	// システムそのもの
	system.NewApp,
)

func di(cfg system.Config) system.App {
	wire.Build(superSet)
	return system.App{}
}
