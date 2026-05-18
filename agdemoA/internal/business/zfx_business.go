package business

import "go.uber.org/fx"

var FxBusinessModule = fx.Module("fx-business-module",
	fx.Provide(
		NewHzwABusinessService,
	),
)
