package svchzw

// import (
// 	smw "ag-core/ag/ag_servicev2"
// 	demob "agdemoB/api/demob"
// 	service "agdemoB/internal/service"
// 	"context"
// 	"fmt"
// 	slog "log/slog"
// 	"sync"
// )

// // DemobProxy 服务代理 - 改进版本
// type DemobProxy struct {
// 	// 核心依赖
// 	impl        *service.DemobImpl
// 	serviceInfo smw.ServiceInfo

// 	// 端点管理
// 	endpoints map[string]smw.Endpoint
// 	mu        sync.RWMutex // 保护endpoints并发访问

// 	// 工具组件
// 	logger *slog.Logger
// }

// // MethodConfig 方法配置信息
// type MethodConfig struct {
// 	Name         string
// 	RequestType  string
// 	ResponseType string
// 	Middlewares  []smw.ServiceMiddlewareProvider
// }

// // NewDemobProxy 创建DemobProxy - 改进版本
// func NewDemobProxy(
// 	service *service.DemobImpl,
// 	mws []smw.ServiceMiddlewareProvider,
// 	opts ...ProxyOption,
// ) demob.Demob {
// 	// 基础服务信息
// 	sinfo := smw.ServiceInfo{
// 		PackageName: "api.demob",
// 		ServiceName: "demob",
// 		HandlerType: demob.Demob(nil),
// 	}

// 	// 创建代理实例
// 	proxy := &DemobProxy{
// 		impl:        service,
// 		endpoints:   make(map[string]smw.Endpoint),
// 		serviceInfo: sinfo,
// 		logger:      slog.Default(),
// 	}

// 	// 应用选项
// 	config := defaultProxyConfig()
// 	// for _, opt := range opts {
// 	// 	opt(&config)
// 	// }

// 	// 注册所有服务方法
// 	proxy.registerEndpoints(mws, config)

// 	return proxy
// }

// // registerEndpoints 注册所有端点
// func (p *DemobProxy) registerEndpoints(
// 	mws []smw.ServiceMiddlewareProvider,
// 	config ProxyConfig,
// ) {
// 	// 定义方法配置
// 	methods := []MethodConfig{
// 		{
// 			Name:         "Calldemob",
// 			RequestType:  "demob.BRequest",
// 			ResponseType: "demob.BReply",
// 			Middlewares:  mws,
// 		},
// 		// 可以轻松添加更多方法
// 		// {
// 		//     Name:         "AnotherMethod",
// 		//     RequestType:  "demob.AnotherRequest",
// 		//     ResponseType: "demob.AnotherReply",
// 		//     Middlewares:  mws,
// 		// },
// 	}

// 	for _, method := range methods {
// 		p.registerEndpoint(method, config)
// 	}
// }

// // registerEndpoint 注册单个端点
// func (p *DemobProxy) registerEndpoint(
// 	method MethodConfig,
// 	config ProxyConfig,
// ) {
// 	// 创建方法级别的调用信息
// 	callInfo := smw.CallInfo{
// 		ServiceInfo: p.serviceInfo,
// 		MethodName:  method.Name,
// 	}

// 	// 构建端点处理链
// 	endpoint := smw.BuilderEndpointChain(
// 		callInfo, // 使用CallInfo传递方法级别信息
// 		method.Middlewares,
// 		p.createHandler(method),
// 	)

// 	p.mu.Lock()
// 	p.endpoints[method.Name] = endpoint
// 	p.mu.Unlock()
// }

// // createHandler 创建方法处理器
// func (p *DemobProxy) createHandler(method MethodConfig) smw.Endpoint {
// 	switch method.Name {
// 	case "Calldemob":
// 		return func(ctx context.Context, req interface{}) (interface{}, error) {
// 			return p.impl.Calldemob(ctx, req.(*demob.BRequest))
// 		}
// 	// 可以轻松扩展其他方法
// 	// case "AnotherMethod":
// 	//     return func(ctx context.Context, req interface{}) (interface{}, error) {
// 	//         return p.impl.AnotherMethod(ctx, req.(*demob.AnotherRequest))
// 	//     }
// 	default:
// 		return func(ctx context.Context, req interface{}) (interface{}, error) {
// 			return nil, fmt.Errorf("handler for method %s not implemented", method.Name)
// 		}
// 	}
// }

// // Calldemob DemobProxy代理方法 - 改进版本
// func (p *DemobProxy) Calldemob(
// 	ctx context.Context,
// 	req *demob.BRequest,
// ) (*demob.BReply, error) {
// 	methodName := "Calldemob"

// 	// 1. 获取端点处理器
// 	p.mu.RLock()
// 	handler, ok := p.endpoints[methodName]
// 	p.mu.RUnlock()

// 	if !ok {
// 		p.logger.Error("endpoint not found",
// 			"method", methodName,
// 			"service", p.serviceInfo.ServiceName)
// 		return nil, fmt.Errorf("endpoint %s not found in service %s",
// 			methodName, p.serviceInfo.ServiceName)
// 	}

// 	// 2. 添加上下文信息
// 	ctx = p.enrichContext(ctx, methodName)

// 	// 3. 执行处理链
// 	res, err := handler(ctx, req)
// 	if err != nil {
// 		p.logger.Error("request handling failed",
// 			"method", methodName,
// 			"service", p.serviceInfo.ServiceName,
// 			"error", err)
// 		return nil, err
// 	}

// 	// 4. 记录成功日志（可选）
// 	if p.logger.Enabled(ctx, slog.LevelDebug) {
// 		p.logger.Debug("request handled successfully",
// 			"method", methodName,
// 			"service", p.serviceInfo.ServiceName)
// 	}

// 	return res.(*demob.BReply), nil
// }

// // enrichContext 丰富上下文信息
// func (p *DemobProxy) enrichContext(ctx context.Context, methodName string) context.Context {
// 	// 添加方法名到上下文
// 	ctx = context.WithValue(ctx, "method", methodName)
// 	ctx = context.WithValue(ctx, "service", p.serviceInfo.ServiceName)
// 	return ctx
// }

// // GetServiceInfo 获取服务信息
// func (p *DemobProxy) GetServiceInfo() smw.ServiceInfo {
// 	return p.serviceInfo
// }

// // GetEndpointCount 获取端点数量（用于监控）
// func (p *DemobProxy) GetEndpointCount() int {
// 	p.mu.RLock()
// 	defer p.mu.RUnlock()
// 	return len(p.endpoints)
// }

// // ProxyOption 代理配置选项
// type ProxyOption func(*ProxyConfig)

// // ProxyConfig 代理配置
// type ProxyConfig struct {
// 	EnableDebugLog bool
// 	CustomLogger   *slog.Logger
// }

// func defaultProxyConfig() ProxyConfig {
// 	return ProxyConfig{
// 		EnableDebugLog: false,
// 		CustomLogger:   slog.Default(),
// 	}
// }

// func WithDebugLog(enabled bool) ProxyOption {
// 	return func(c *ProxyConfig) {
// 		c.EnableDebugLog = enabled
// 	}
// }

// func WithLogger(logger *slog.Logger) ProxyOption {
// 	return func(c *ProxyConfig) {
// 		c.CustomLogger = logger
// 	}
// }
