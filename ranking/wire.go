//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lazywoo/mercury/pkg/wego"
	"github.com/lazywoo/mercury/ranking/grpc"
	"github.com/lazywoo/mercury/ranking/ioc"
	"github.com/lazywoo/mercury/ranking/repository"
	"github.com/lazywoo/mercury/ranking/repository/cache"
	"github.com/lazywoo/mercury/ranking/service"
)

var thirdProviderSet = wire.NewSet(
	ioc.InitLogger,
	ioc.InitRedis,
	ioc.InitEtcdClient,
	ioc.InitArticleRpcClient,
	ioc.InitInteractiveRpcClient,
)

var svcProviderSet = wire.NewSet(
	service.NewBatchRankingService,
	repository.NewRankingCachedRepository,
	cache.NewRankingLocalCache,
	cache.NewRankingRedisCache,
)

var cronProviderSet = wire.NewSet(
	ioc.InitTasks,
	ioc.InitRankingJob,
	ioc.InitRLockClient,
)

func InitAPP() *wego.App {
	wire.Build(
		thirdProviderSet,
		svcProviderSet,
		cronProviderSet,
		grpc.NewRankingServiceServer,
		ioc.InitGRPCxServer,
		wire.Struct(new(wego.App), "GRPCServer", "Cron"),
	)
	return new(wego.App)
}
