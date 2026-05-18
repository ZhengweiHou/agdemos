package repository

import (
	"agdemoC/internal/repository/dao"

	"go.uber.org/fx"
)

var FxRepositoryModule = fx.Module("fx-repository-module",
	fx.Provide(
		dao.NewStudentDao,
	),
)
