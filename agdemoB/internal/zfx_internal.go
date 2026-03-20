package internal

import (
	"agdemoB/internal/hzwdev"

	"go.uber.org/fx"
)

var FxInternalModule = fx.Module("fx-internal-module",
	// headdev.FxHeadDevModule,
	fx.Provide(
		hzwdev.FxHertzHzwCustMiddlewareProvider,
	),
)
