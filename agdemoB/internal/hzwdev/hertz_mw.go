package hzwdev

import (
	"ag-core/contribute/aghertz/server"
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
)

func HertzHzwCustMiddleware() app.HandlerFunc {
	return func(ctx context.Context, r *app.RequestContext) {
		host := r.Host()
		fmt.Printf("host: %s\n", host)
		r.Next(ctx)
	}
}

var FxHertzHzwCustMiddlewareProvider = server.NewFxServerMiddlewareProvider(
	HertzHzwCustMiddleware,
)
