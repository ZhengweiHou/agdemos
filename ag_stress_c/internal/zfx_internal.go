package internal

import (
	"ag_stress_c/internal/adpgen"
	"ag_stress_c/internal/svcgen"

	"go.uber.org/fx"
)

var FxInternalModule = fx.Module("fx-internal-module",
	svcgen.FxServiceWithProxyModule(),
	adpgen.FxAdapterModule(),
)
