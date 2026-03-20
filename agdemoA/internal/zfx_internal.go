package internal

import (
	"agdemoA/internal/business"
	"agdemoA/internal/clients"
	"agdemoA/internal/config"
	"agdemoA/internal/gatewayimpl"

	"go.uber.org/fx"
)

var FxInternalModule = fx.Module("fx-internal-module",
	business.FxBusinessModule,
	gatewayimpl.FxGatewayImplModule,
	clients.FxClientModule,
	config.FxConfigModule,
	// headdev.FxHeadDevModule
)
