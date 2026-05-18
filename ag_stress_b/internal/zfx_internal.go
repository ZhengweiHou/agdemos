package internal

import (
	"ag_stress_b/internal/adpgen"
	"ag_stress_b/internal/clients"
	"ag_stress_b/internal/domain"
	"ag_stress_b/internal/svcgen"

	"go.uber.org/fx"
)

var FxInternalModule = fx.Module("fx-internal-module",
	svcgen.FxServiceWithProxyModule(),
	adpgen.FxAdapterModule(),

	domain.FxDomainModule,
	clients.FxClientModule,
)
