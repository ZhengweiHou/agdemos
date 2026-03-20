package gatewayimpl

import "go.uber.org/fx"

var FxGatewayImplModule = fx.Module("fx-gateway-impl-module",
	fx.Provide(
		NewCallDemobGatewayImpl,
	),
)
