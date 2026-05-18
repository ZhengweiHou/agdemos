package svchzw

import (
	ag_service "ag-core/ag/ag_service"
	smw "ag-core/ag/ag_service"
	demob "agdemoB/api/demob"
	service "agdemoB/internal/service"
	"context"
	"fmt"
	"log/slog"
)

// var DemobInfoHelper demebInfoHelper

// type demebInfoHelper int

// func (d demebInfoHelper) ServiceInfo() ag_service.ServiceInfo {
// 	return DemobServiceInfo
// }

// func (d demebInfoHelper) Calldemob() ag_service.CallInfo {
// 	return Demob_Calldemob_CallInfo
// }

// 服务信息,接口级别
//
//	var DemobServiceInfo = ag_service.ServiceInfo{
//		PackageName: "api.demob",
//		ServiceName: "demob",
//		HandlerType: demob.Demob(nil),
//	}
var DemobServiceInfo = ag_service.NewServiceInfo(
	"api.demob",
	"demob",
	demob.Demob(nil),
)

// 方法级别的调用信息,servicename+methodname+CallInfo // TODO 需要到包级别的区分吗?
//
//	var Demob_Calldemob_CallInfo = &ag_service.CallInfo{
//		ServiceInfo: DemobServiceInfo,
//		CallName:    "Calldemob",
//		Extra: map[string]interface{}{
//			"method_name": "Calldemob",
//		},
//	}
var Demob_Calldemob_CallInfo = ag_service.NewCallInfo(
	DemobServiceInfo,
	"Calldemob",
	false,
	false,
)

var DemobCallInfos = map[string]*ag_service.CallInfo{
	"Calldemob": Demob_Calldemob_CallInfo,
}

// DemobProxy 服务代理
type DemobProxy_hzw struct {
	ag_service.AgServiceProxyBase

	impl demob.Demob

	// ServiceInfo ag_service.ServiceInfo
	// MethodInfos map[string]ag_service.CallInfo

	// mu        sync.RWMutex // 保护endpoints并发访问
	// endpoints map[string]ag_service.Endpoint
}

// NewDemobProxy 创建DemobProxy
func NewDemobProxy_hzw(
	impl *service.DemobImpl,
	agServiceBuilder ag_service.AgServiceBuilder,
	mws []ag_service.MiddlewareProvider,
) (proxy demob.Demob, err error) {
	p := &DemobProxy_hzw{
		impl: impl, // 实际代理目标
		AgServiceProxyBase: ag_service.NewAgServiceProxyBase(
			DemobServiceInfo,
			DemobCallInfos,
		),
		// AgServiceProxyBase: ag_service.DefaultAgServiceProxyBase(),
		// endpoints:   make(map[string]ag_service.Endpoint),
		// ServiceInfo: DemobServiceInfo,
		// MethodInfos: DemobCallInfos,
	}

	for _, cif := range p.CallInfos {
		if cif.IsClientStreaming() || cif.IsServerStreaming() {
			slog.Warn("agServiceProxy buildskip streaming method", "call_name", cif.CallName)
			continue
		}
		// 根据方法信息构建Endpoints
		// endpoint, err := ag_service.BuildEndpointChain(
		endpoint, err := agServiceBuilder.BuildEndpointChain(
			cif,
			mws,
			p.getOriginalHandler(cif),
		)
		if err != nil {
			return nil, err
		}
		// 注册endpoint
		// p.registerEndpoint(cif.CallName, endpoint)
		p.RegisterEndpoint(cif.CallName(), endpoint)
	}

	return p, nil
}

func (p *DemobProxy_hzw) Calldemob(ctx context.Context, req *demob.BRequest) (*demob.BReply, error) {
	// endpoint := p.getEndpoint("Calldemob")
	endpoint := p.GetEndpoint("Calldemob")
	// 调用endpoint
	resp, err := endpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*demob.BReply), nil
}

// // registerEndpoint 注册单个端点
// func (p *DemobProxy_hzw) registerEndpoint(callName string, endpoint ag_service.Endpoint) {
// 	p.mu.Lock()
// 	p.endpoints[callName] = endpoint
// 	p.mu.Unlock()
// }

// func (p *DemobProxy_hzw) getEndpoint(callName string) ag_service.Endpoint {
// 	p.mu.RLock()
// 	defer p.mu.RUnlock()
// 	return p.endpoints[callName]
// }

// 获取原始方法处理器
func (p *DemobProxy_hzw) getOriginalHandler(cif *ag_service.CallInfo) smw.Endpoint {
	switch cif.CallName() {
	case "Calldemob":
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return p.impl.Calldemob(ctx, req.(*demob.BRequest))
		}
	// 可以轻松扩展其他方法
	// case "AnotherMethod":
	//     return func(ctx context.Context, req interface{}) (interface{}, error) {
	//         return p.impl.AnotherMethod(ctx, req.(*demob.AnotherRequest))
	//     }
	default:
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return nil, fmt.Errorf("handler for method %s not implemented", cif.CallName)
		}
	}
}

// func (p *DemobProxy_hzw) ClientSideStreaming(stream demob.Demob_ClientSideStreamingServer) error {
// 	// TODO
// 	return nil
// }

// func (p *DemobProxy_hzw) ServerSideStreaming(req *demob.BRequest, stream demob.Demob_ServerSideStreamingServer) error {
// 	// TODO
// 	return nil
// }

// func (p *DemobProxy_hzw) BidiSideStreaming(stream demob.Demob_BidiSideStreamingServer) error {
// 	// TODO
// 	return nil
// }

func (c *DemobProxy_hzw) Hello(ctx context.Context, req *demob.BRequest) (*demob.BReply, error) {
	var resp *demob.BReply
	resp = &demob.BReply{}
	// TODO 添加业务处理逻辑

	return resp, nil
}
