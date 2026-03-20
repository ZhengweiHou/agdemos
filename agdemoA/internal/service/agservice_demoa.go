// TODO 这里需要根据业务场景，自行修改实现

package service

import (
	demoa "agdemoA/api/demoa"
	"agdemoA/internal/business"
	"agdemoA/internal/business/model"
	"context"
)

// DemoAImpl defines the service implementation for DemoAImpl.
type DemoAImpl struct {
	HzwAService *business.HzwABusinessService
}

// NewDemoAImpl creates and returns a new DemoAImpl instance.
// @param TODO inject dependencies
// @return *DemoAImpl
func NewDemoAImpl(hzwAService *business.HzwABusinessService) *DemoAImpl {
	return &DemoAImpl{
		HzwAService: hzwAService,
	}
}

// NewDemoAImplI new DemoAImpl instance, and return interface demoa.DemoA.
// @param TODO inject dependencies
// @return demoa.DemoA
func NewDemoAImplI(hzwAService *business.HzwABusinessService) demoa.DemoA {
	return NewDemoAImpl(hzwAService)
}

// Calldemoa TODO:DESCRIBE
func (c *DemoAImpl) Calldemoa(ctx context.Context, req *demoa.ARequest) (*demoa.AReply, error) {
	var resp *demoa.AReply
	resp = &demoa.AReply{}

	// 调用业务服务
	mreq := &model.HzwModelReq{
		Name: req.Name,
	}
	mresp, err := c.HzwAService.DoSomeThing(ctx, mreq)
	if err != nil {
		return nil, err
	}

	resp.Msg = mresp.Msg
	return resp, nil
}
