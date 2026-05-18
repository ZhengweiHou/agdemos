// TODO 这里需要根据业务场景，自行修改实现

package service

import (
	stressb "ag_stress_b/api/stressb"
	"ag_stress_b/internal/domain"
	"context"
)

// BServiceImpl defines the service implementation for BServiceImpl.
type BServiceImpl struct {
	bDomainSvc *domain.BDomainService
}

// NewBServiceImpl creates and returns a new BServiceImpl instance.
func NewBServiceImpl(
	bDomainSvc *domain.BDomainService,
) *BServiceImpl {
	return &BServiceImpl{
		bDomainSvc: bDomainSvc,
	}
}

// EchoB TODO:DESCRIBE
func (c *BServiceImpl) EchoB(ctx context.Context, req *stressb.EchoBRequest) (*stressb.EchoBResponse, error) {
	resp := &stressb.EchoBResponse{}

	msg, err := c.bDomainSvc.DoB(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	resp.Msg = msg
	return resp, nil
}
