package clients

import "go.uber.org/fx"

var FxClientModule = fx.Module("fx-clients-module",
	fx.Provide(
		NewDemoBHertzClient,
		NewDemoKitexClient,
	),
)
