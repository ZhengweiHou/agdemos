package headdev

import (
	"ag-core/ag/ag_common/agmetadata"
	"context"
	"fmt"
	"log/slog"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
)

func KitexServerHeadMiddlewar(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		if !isGRPC(ri) {
			return next(ctx, req, resp)
		}
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return fmt.Errorf("kitex from incoming context not found")
		}
		slog.Debug("=== TEST kitex head middleware", "md", md)

		rctx, err := agmetadata.ParseMdToContext(ctx, func(key string) ([]string, bool) {
			tmph := md.Get(key)
			if len(tmph) == 1 {
				return []string{tmph[0]}, true
			} else if len(tmph) > 1 {
				slog.Warn("kitex head key has multiple values, metadata will use the first value", "key", key)
				return []string{tmph[0]}, true
			} else {
				return nil, false
			}
		})
		if err != nil {
			return err
		}
		err = next(rctx, req, resp)

		return err
	}
}

func KitexClientHeadMiddlewar(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		// 从上下文提取头信息
		md := agmetadata.GetMdFromContext(ctx)
		var kmd []string
		for k, v := range md { // TODO 是否所有的元数据都透传
			kmd = append(kmd, k, v)
		}

		rctx := ctx
		if len(kmd) > 0 {
			rctx = metadata.AppendToOutgoingContext(ctx, kmd...)
		}

		return next(rctx, req, resp)
	}
}

func NewKitexServerHeadMiddleware() endpoint.Middleware {
	return KitexServerHeadMiddlewar
}

func NewKitexClientHeadMiddleware() endpoint.Middleware {
	return KitexClientHeadMiddlewar
}
func isGRPC(ri rpcinfo.RPCInfo) bool {
	return ri.Config().TransportProtocol()&transport.GRPC == transport.GRPC
}
