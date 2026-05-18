package domain

import (
	"ag-core/ag/ag_ext/future"
	bhsvc "ag_stress_a/internal/adpgen/hertz/bservice"
	bksvc "ag_stress_a/internal/adpgen/kitex/bservice"
	stressb "ag_stress_b/api/stressb"
	"context"
	"log/slog"
	"strings"

	"github.com/panjf2000/ants/v2"
)

type ADomainService struct {
	bhttpcli bhsvc.BServiceHertzClient
	bgrpccli bksvc.Client
	antpool  *ants.Pool
}

func NewADomainService(
	bhttpcli bhsvc.BServiceHertzClient,
	bgrpccli bksvc.Client,
) *ADomainService {

	options := ants.Options{
		// PanicHandler: func(a any) {
		// 	slog.Error("goroutine pool panic", "err", a)
		// },
	}
	antpool, _ := ants.NewPool(3, ants.WithOptions(options))

	return &ADomainService{
		bhttpcli: bhttpcli,
		bgrpccli: bgrpccli,
		antpool:  antpool,
	}
}

func (s *ADomainService) DoA(ctx context.Context, name string) (string, error) {
	slog.Info("[I] ag_stress_a DoA")
	defer func() {
		slog.Info("[O] ag_stress_a DoA end")
	}()

	echoResp, err := s.echoB(ctx, name)
	if err != nil {
		return "", err
	}

	return echoResp, nil
}

func (s *ADomainService) echoB(ctx context.Context, name string) (string, error) {
	req := &stressb.EchoBRequest{Name: name}
	var (
		resp *stressb.EchoBResponse
		err  error
	)

	msg := "default_msg"

	if strings.Contains(name, "bgrpc") {
		slog.Info("echoB with grpc")

		// resp, err = s.bgrpccli.EchoB(ctx, req)
		fut := future.NewFuture(func() (*stressb.EchoBResponse, error) {
			return s.bgrpccli.EchoB(ctx, req)
		})
		// fut := s.bgrpccli.EchoBAsync(ctx, req)
		resp, err = fut.Await(ctx)
		if err != nil {
			return "", err
		}
	} else if strings.Contains(name, "anil") {
		gpanic1(s)
		gpanic2(s)
		gpanic3(s)
		gpanic4(s)
	} else {
		slog.Info("echoB with http")
		// resp, err = s.bhttpcli.EchoB(ctx, req)
		fut := future.NewFuture(func() (*stressb.EchoBResponse, error) {
			return s.bhttpcli.EchoB(ctx, req)
		})
		resp, err = fut.Await(ctx)
		if err != nil {
			return "", err
		}
	}

	if resp != nil {
		msg = resp.Msg
	}

	return msg, nil
}

func gpanic1(s *ADomainService) {
	_ = future.NewFutureFunc(func() (interface{}, error) {
		panic("gpanic1")
	})
}

func gpanic2(s *ADomainService) {
	future.FutureCall(func() (interface{}, error) {
		panic("gpanic2")
	}, func(res interface{}, err error) {
	})
}

func gpanic3(s *ADomainService) {
	_ = future.NewFuture(func() (interface{}, error) {
		panic("gpanic3")
	})
}

func gpanic4(s *ADomainService) {
	go func() {
		s.antpool.Submit(func() {
			panic("gpanic4")
		})
	}()
}
