package gateway

import (
	"context"
)

type CHelloReq struct {
	Name string `json:"name"`
}

type CHelloReply struct {
	Msg string `json:"msg"`
}

type IDemoCGateway interface {
	Hello(ctx context.Context, req *CHelloReq) (resp *CHelloReply, err error)
}
