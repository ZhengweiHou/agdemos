package domain

import "go.uber.org/fx"

var FxDomainModule = fx.Module("fx-domain-module",
	fx.Provide(
		NewADomainService,
	),
)
