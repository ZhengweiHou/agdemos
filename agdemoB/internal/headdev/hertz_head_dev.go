package headdev

import (
	"ag-core/ag/ag_common/agmetadata"
	"ag-core/contribute/aghertz/server"
	"context"
	"log/slog"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
)

func HertzServerHeadMiddleware(ctx context.Context, r *app.RequestContext) {
	slog.Info("hertz server head middleware")

	rctx, err := agmetadata.ParseMdToContext(ctx, func(key string) ([]string, bool) {
		tmph := r.GetHeader(key)
		if len(tmph) > 0 {
			return []string{string(tmph)}, true
		} else {
			return nil, false
		}
	})
	if err != nil {
		slog.Error("parse md to context error", "err", err)
		return
	}

	r.Next(rctx)
}

// HertzClientHeadMiddleware  Hertz 客户端头信息中间件
func HertzClientHeadMiddleware(endpoint client.Endpoint) client.Endpoint {
	return func(ctx context.Context, req *protocol.Request, resp *protocol.Response) (err error) {

		// 从上下文提取头信息
		md := agmetadata.GetMdFromContext(ctx)
		for key, value := range md { // TODO 是否所有的元数据都透传
			req.Header.Set(key, value)
		}

		return endpoint(ctx, req, resp)
	}
}

func NewHertzServerHeadMiddleware() server.Middleware {
	return HertzServerHeadMiddleware
}

func NewHertzClientHeadMiddleware() client.Middleware {
	return HertzClientHeadMiddleware
}
