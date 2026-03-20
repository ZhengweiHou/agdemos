package gatewayimpl

import (
	// demobapi "agdemoA/api/demob"
	hdemob "agdemoA/internal/adpgen/hertz/demob"
	kdemob "agdemoA/internal/adpgen/kitex/demob"
	"agdemoA/internal/business/gateway"
	demobapi "agdemoB/api/demob"
	"context"
	"log/slog"
	"strings"
)

type CallDemobGatewayImpl struct {
	DemobHertzClient hdemob.DemobHertzClient
	DemobKitexClient kdemob.Client
}

func NewCallDemobGatewayImpl(bclient hdemob.DemobHertzClient, kclient kdemob.Client) gateway.ICallDemobGateway {
	return &CallDemobGatewayImpl{
		DemobHertzClient: bclient,
		DemobKitexClient: kclient,
	}
}

func (g *CallDemobGatewayImpl) CallDemob(ctx context.Context, req *gateway.CallDemobReq) (resp *gateway.CallDemobResp, err error) {
	name := req.Name
	resp = &gateway.CallDemobResp{}
	breq := &demobapi.BRequest{
		Name: req.Name,
	}

	bresp := &demobapi.BReply{}

	slog.Info("[S] agdemoA CallDemobGatewayImpl CallDemob")
	if strings.HasPrefix(name, "grpc") {
		bresp, err = g.DemobKitexClient.Calldemob(
			ctx,
			breq,
			// callopt.WithFallback(fallback.NewFallbackPolicy( // Call维度配置
			// 	func(ctx context.Context, args utils.KitexArgs, result utils.KitexResult, err error) (fbErr error) {
			// 		return err
			// 	})),
		)
	} else {
		bresp, err = g.DemobHertzClient.Calldemob(ctx, breq)
	}
	slog.Info("[F] agdemoA CallDemobGatewayImpl CallDemob")

	if err != nil {
		return nil, err
	}

	resp.Msg = bresp.Msg

	return resp, nil
}
