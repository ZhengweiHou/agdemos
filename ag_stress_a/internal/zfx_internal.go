package internal

import (
	"ag_stress_a/internal/adpgen"
	"ag_stress_a/internal/clients"
	"ag_stress_a/internal/domain"
	"ag_stress_a/internal/svcgen"

	"go.uber.org/fx"
)

var FxInternalModule = fx.Module("fx-internal-module",
	svcgen.FxServiceWithProxyModule(),
	adpgen.FxAdapterModule(),

	domain.FxDomainModule,
	clients.FxClientModule,

	// fx.Provide(
	// 	resolver.BuildKitexHzwResolver,
	// ),

	// resolver.HzwResolverConditional,

)
