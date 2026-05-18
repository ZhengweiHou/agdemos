package domain

import (
	chsvc "ag_stress_b/internal/adpgen/hertz/cservice"
	cksvc "ag_stress_b/internal/adpgen/kitex/cservice"
	"ag_stress_c/api/stressc"
	"context"
	"log/slog"
	"strings"
)

type BDomainService struct {
	chttpcli chsvc.CServiceHertzClient
	cgrpccli cksvc.Client
}

func NewADomainService(
	chttpcli chsvc.CServiceHertzClient,
	cgrpccli cksvc.Client,
) *BDomainService {
	return &BDomainService{
		chttpcli: chttpcli,
		cgrpccli: cgrpccli,
	}
}

func (s *BDomainService) DoB(ctx context.Context, name string) (string, error) {
	slog.Info("[I] ag_stress_b DoB")
	defer func() {
		slog.Info("[O] ag_stress_b DoB end")
	}()

	echoResp, err := s.echoC(ctx, name)
	if err != nil {
		return "", err
	}

	return echoResp, nil
}

func (s *BDomainService) echoC(ctx context.Context, name string) (string, error) {
	req := &stressc.EchoCRequest{Name: name}
	var (
		resp *stressc.EchoCResponse
		err  error
	)

	if strings.Contains(name, "cgrpc") {
		slog.Info("echoC with grpc")
		resp, err = s.cgrpccli.EchoC(ctx, req)
		if err != nil {
			return "", err
		}

	} else {
		slog.Info("echoC with http")
		resp, err = s.chttpcli.EchoC(ctx, req)
		if err != nil {
			return "", err
		}
	}

	msg := resp.Msg
	return msg, nil
}
