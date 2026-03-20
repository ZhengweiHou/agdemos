package headdev

import (
	hclient "ag-core/contribute/aghertz/client"
	hserver "ag-core/contribute/aghertz/server"

	kclient "ag-core/contribute/agkitex/client"
	kserver "ag-core/contribute/agkitex/server"

	"go.uber.org/fx"
)

var FxHeadDevModule = fx.Module("fx-headdev",
	fx.Provide(
		hserver.NewFxServerMiddlewareProvider(
			NewHertzServerHeadMiddleware,
		),
		hclient.NewFxClientMiddlewareProvider(
			NewHertzClientHeadMiddleware,
		),

		kserver.NewFxServerMiddlewareProvider(
			NewKitexServerHeadMiddleware,
		),
		kclient.NewFxClientMiddlewareProvider(
			NewKitexClientHeadMiddleware,
		),
	),
)
