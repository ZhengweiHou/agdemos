package gatewayimpl

import (
	// demobapi "agdemoA/api/demob"

	helloapi "agdemoA/api/hello"
	hhello "agdemoA/internal/adpgen/hertz/hello"
	khello "agdemoA/internal/adpgen/kitex/hello"
	"agdemoA/internal/business/gateway"
	"context"
	"log/slog"
	"strings"
)

type CallCGatewayImpl struct {
	khellocli khello.Client
	hhellocli hhello.HelloHertzClient
}

func NewCallCGatewayImpl(
	kclient khello.Client,
	hclient hhello.HelloHertzClient,
) gateway.IDemoCGateway {
	return &CallCGatewayImpl{
		khellocli: kclient,
		hhellocli: hclient,
	}
}

func (g *CallCGatewayImpl) Hello(ctx context.Context, req *gateway.CHelloReq) (resp *gateway.CHelloReply, err error) {
	resp = &gateway.CHelloReply{}
	breq := &helloapi.HelloRequest{
		Name: req.Name,
	}

	bresp := &helloapi.HelloReply{}

	name := req.Name

	slog.Info("[S] agdemoA CallDemobGatewayImpl CallDemob")
	if strings.Contains(name, "grpc") {
		slog.Info("使用grpc调用C")
		bresp, err = g.khellocli.Pinghello(
			ctx,
			breq,
		)
	} else {
		slog.Info("使用http调用C")
		bresp, err = g.hhellocli.Pinghello(
			ctx,
			breq,
		)
	}
	slog.Info("[F] agdemoA CallDemobGatewayImpl CallDemob")

	if err != nil {
		return nil, err
	}

	resp.Msg = bresp.Msg

	return resp, nil
}
