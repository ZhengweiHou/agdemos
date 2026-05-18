package svchzw

import (
	"ag-core/ag/ag_service"
	"context"
	"fmt"
)

func HzwCallinfoOpt(ci *ag_service.CallInfo) error {
	fmt.Printf("====== callinfoopt: %v\n", ci)
	ci.Extra["HzwCallinfoOpt"] = true
	return nil
}

func HzwMWFunc(next ag_service.Endpoint) ag_service.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		cif := ag_service.GetCallInfoFromContext(ctx)

		fmt.Printf("====== HzwMWFunc HzwCallinfoOpt: %v\n", cif.Extra["HzwCallinfoOpt"])
		htype := cif.ServiceInfo().HandlerType()
		fmt.Printf("====== HzwMWFunc HandlerType: %T\n", htype)

		return next(ctx, req)
	}
}

var HzwCIFOptFx = ag_service.NewFxAgCallInfoOpt(
	func() ag_service.CallInfoOpt {
		return HzwCallinfoOpt
	},
)

var HzwMWFx = ag_service.NewFxAgGlobalMiddleware(
	func() ag_service.MiddlewareFunc {
		return HzwMWFunc
	},
)
