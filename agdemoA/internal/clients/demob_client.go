package clients

import (
	"ag-core/contribute/aghertz/aghertzclient"
	hdemob "agdemoA/internal/adpgen/hertz/demob"
	kdemob "agdemoA/internal/adpgen/kitex/demob"
	"context"
	"log/slog"

	agkclient "ag-core/contribute/agkitex/client"

	"github.com/cloudwego/kitex/client"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"

	hertzclient "github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
	kslog "github.com/kitex-contrib/obs-opentelemetry/logging/slog"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func NewDemoBHertzClient(cli *hertzclient.Client) hdemob.DemobHertzClient {

	// TODO 获取服务端点名配置：
	demobEndpoint := "agdemob_hertz"
	// demobEndpoint := "127.0.0.1:9997"

	bnfshcli := hdemob.NewDemobHertzClient(
		cli,
		aghertzclient.WithEndpoint(demobEndpoint),
		aghertzclient.WithSD(true), // 开启服务发现
	)
	return bnfshcli
}

func NewDemoKitexClient(suite *agkclient.KitexClientSuite) (kdemob.Client, error) {

	// TODO 获取服务端点名配置：
	demobEndpoint := "agdemob_kitex"

	opts := []kclient.Option{}
	opts = append(opts,
		hzwFallbackPolice(),
		hzwFailureRetry(),
		hzwClientTracer(),

		client.WithSuite(tracing.NewClientSuite()),
		// client.WithSuite(tracing.NewGRPCClientSuite()), // NewGRPCClientSuite已弃用
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "agdemoa"}),
	)

	// bnfshcli, err := kdemob.NewClient(demobEndpoint)
	bnfshcli, err := kdemob.NewClientWithSuite(
		demobEndpoint,
		suite,
		opts...,
	)
	return bnfshcli, err
}

// 降级样例
func hzwFallbackPolice() kclient.Option {
	return kclient.WithFallback( // 降级
		fallback.NewFallbackPolicy(func(ctx context.Context, args utils.KitexArgs, result utils.KitexResult, err error) (fbErr error) {
			slog.Info("==kitex fallback")

			// result.SetSuccess()
			return
		}),
		// fallback.TimeoutAndCBFallback(
		// 	fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
		// 		// 打印请求及响应的类型
		// 		slog.Info(fmt.Sprintf("req type: %T, resp type: %T", req, resp))
		// 		slog.Error("==kitex fallback", "err", err)
		// 		// do something
		// 		return resp, err
		// 	}),
		// ),
	)
}

func hzwFailureRetry() kclient.Option {
	// import "github.com/cloudwego/kitex/pkg/retry"
	fp := retry.NewFailurePolicy()

	fp.WithMaxRetryTimes(3) // 配置最多重试3次

	return kclient.WithFailureRetry( // 失败重试
		fp,
	)
}

func hzwClientTracer() kclient.Option {
	return kclient.WithTracer( // 开启tracing
		&hzwtracer{},
	)
}

type hzwtracer struct {
}

func (t *hzwtracer) Start(ctx context.Context) context.Context {
	slog.Info("==tracer start")
	return ctx
}
func (t *hzwtracer) Finish(ctx context.Context) {
	slog.Info("==tracer finish")
}

func init() {
	// 配置 Provider
	// p := provider.NewOpenTelemetryProvider(
	// 	provider.WithServiceName("agdemoa"),
	// 	provider.WithExportEndpoint("localhost:4317"), // OTLP gRPC
	// 	// provider.WithExportEndpoint("localhost:4318"),      // OTLP HTTP（可选）
	// 	provider.WithInsecure(),
	// 	// provider.WithEnableMetrics(true), // 启用指标
	// 	provider.WithEnableTracing(true), // 启用追踪
	// 	provider.WithResourceAttributes(
	// 		[]attribute.KeyValue{
	// 			{Key: "environment", Value: attribute.StringValue("production")},
	// 			{Key: "version", Value: attribute.StringValue("1.0.0")},
	// 		},
	// 	),
	// )
	// fmt.Sprintf("==otel provider: %v", p)

	// 设置 Kitex 日志集成（可选）
	klog.SetLogger(kslog.NewLogger())
}
