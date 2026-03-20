package repository

import (
	"ag-core/contribute/agdb/agdao"
	"agdemoB/internal/repository/dao"
	"context"
	"log/slog"

	"go.uber.org/fx"
)

var FxAgdemobRepositoryModule = fx.Module(
	"fx_agdemob_repository",
	fx.Provide(
		dao.NewStudentDao,
		// daohzw.NewStudentDao,

		agdao.NewFxAgTbInfoOpt(
			func() agdao.TbInfoOpt {
				return agdao.WithTbNameStrategy(func(ctx context.Context, info *agdao.TableInfo) string {
					slog.Info("=======自定义的表名策略 TEST====")
					return info.TableName
				})
			},
		),
	),
)

// var transMw = fx.Module(
// 	"fx_aic_gormdb",
// 	fx.Provide(
// 	// gormdb.NewTmMiddlewareContext,
// 	// ag_service.NewServiceGlobalMiddleware(
// 	//     func(mv *gormdb.TmMiddlewareContext){

// 	//     }
// 	// ),
// 	),
// )
