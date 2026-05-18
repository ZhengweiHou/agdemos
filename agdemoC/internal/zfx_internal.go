package internal

import (
	"agdemoC/internal/adpgen"
	"agdemoC/internal/repository"
	"agdemoC/internal/svcgen"

	"go.uber.org/fx"
)

var FxInternalModule = fx.Module("fx-internal-module",
	// fx.Provide(
	// 适配层(代码生成)
	adpgen.FxAdapterModule(),
	// 业务层实现(代码生成)
	svcgen.FxServiceWithProxyModule(),
	// 数据库访问层
	repository.FxRepositoryModule,
	// ),
)
