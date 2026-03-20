package business

import (
	"ag-core/ag/ag_common/agmetadata"
	"agdemoA/internal/business/gateway"
	"agdemoA/internal/business/model"
	"context"
	"fmt"
	"log/slog"
)

type HzwABusinessService struct {
	CallDemobService gateway.ICallDemobGateway
}

func NewHzwABusinessService(callDemobService gateway.ICallDemobGateway) *HzwABusinessService {
	return &HzwABusinessService{
		CallDemobService: callDemobService,
	}
}

func (s *HzwABusinessService) DoSomeThing(ctx context.Context, req *model.HzwModelReq) (resp *model.HzwModelResp, err error) {
	slog.Info("[I] agdemoA HzwABusinessService DoSomeThing")
	resp = &model.HzwModelResp{}

	md := agmetadata.GetMdFromContext(ctx)
	slog.Info(fmt.Sprintf("md: %v", md))

	// 调用Demob服务
	callBReq := &gateway.CallDemobReq{
		Name: req.Name,
	}
	callBResp, err := s.CallDemobService.CallDemob(ctx, callBReq)
	if err != nil {
		slog.Error("[E] agdemoA HzwABusinessService DoSomeThing CallDemobService.CallDemob err", "err", err)
		return nil, err
	}

	resp.Msg = callBResp.Msg
	slog.Info("[O] agdemoA HzwABusinessService DoSomeThing")
	return resp, nil
}
