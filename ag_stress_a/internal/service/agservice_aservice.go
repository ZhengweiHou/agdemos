// TODO 这里需要根据业务场景，自行修改实现

package service

import (
	stressa "ag_stress_a/api/stressa"
	"ag_stress_a/internal/domain"
	"context"
)

// AServiceImpl defines the service implementation for AServiceImpl.
type AServiceImpl struct {
	aDomainSvc *domain.ADomainService
}

// NewAServiceImpl creates and returns a new AServiceImpl instance.
// @param TODO inject dependencies
// @return *AServiceImpl
func NewAServiceImpl(
	aDomainSvc *domain.ADomainService,
) *AServiceImpl {
	return &AServiceImpl{
		aDomainSvc: aDomainSvc,
	}
}

// EchoA TODO:DESCRIBE
func (c *AServiceImpl) EchoA(ctx context.Context, req *stressa.EchoARequest) (*stressa.EchoAResponse, error) {
	var resp *stressa.EchoAResponse
	resp = &stressa.EchoAResponse{}

	msg, err := c.aDomainSvc.DoA(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	resp.Msg = msg
	return resp, nil
}
