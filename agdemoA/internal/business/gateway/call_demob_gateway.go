package gateway

import (
	"context"
)

type CallDemobReq struct {
	Name string `json:"name"`
}

type CallDemobResp struct {
	Msg string `json:"msg"`
}

type ICallDemobGateway interface {
	CallDemob(ctx context.Context, req *CallDemobReq) (resp *CallDemobResp, err error)
}
