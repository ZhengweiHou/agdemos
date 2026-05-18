// TODO 这里需要根据业务场景，自行修改实现

package service

import (
	stressc "ag_stress_c/api/stressc"
	"context"
)

// CServiceImpl defines the service implementation for CServiceImpl.
type CServiceImpl struct {
	// TODO 这里需要根据业务场景，添加依赖
}

// NewCServiceImpl creates and returns a new CServiceImpl instance.
// @param TODO inject dependencies
// @return *CServiceImpl
func NewCServiceImpl() *CServiceImpl {
	return &CServiceImpl{}
}

// EchoC TODO:DESCRIBE
func (c *CServiceImpl) EchoC(ctx context.Context, req *stressc.EchoCRequest) (*stressc.EchoCResponse, error) {
	resp := &stressc.EchoCResponse{
		Msg: "echo c " + req.Name,
	}
	// TODO 添加业务处理逻辑

	// time.Sleep(time.Second * 3)

	return resp, nil
}
