package clients

import (
	hhello "agdemoA/internal/adpgen/hertz/hello"
	khello "agdemoA/internal/adpgen/kitex/hello"
	"fmt"

	"ag-core/contribute/aghertz/aghertzclient"
	agkclient "ag-core/contribute/agkitex/client"

	hclient "github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/kitex/client"
	kclient "github.com/cloudwego/kitex/client"

	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func NewHelloHertzClient(cli *hclient.Client) hhello.HelloHertzClient {
	// 获取服务端点名配置：
	// demobEndpoint := "agdemoc_hertz"
	demobEndpoint := "127.0.0.1:1111"
	// demobEndpoint := "127.0.0.1:9997"

	bnfshcli := hhello.NewHelloHertzClient(
		cli,
		aghertzclient.WithEndpoint(demobEndpoint),
		// aghertzclient.WithSD(true), // 开启服务发现
	)
	return bnfshcli
}

func NewHelloKitexClient(suite *agkclient.KitexClientSuite) (khello.Client, error) {
	// 获取服务端点名配置：
	// demobEndpoint := "agdemoc_kitex"
	demobEndpoint := "hzw-cloud_gateway-v3"

	opts := []kclient.Option{}
	opts = append(opts,
		client.WithSuite(tracing.NewClientSuite()),
		// client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "agdemoa"}),
		client.WithHostPorts(fmt.Sprintf("127.0.0.1:%d", 1111)), // 直连
	)

	// bnfshcli, err := kdemob.NewClient(demobEndpoint)
	hellocli, err := khello.NewClientWithSuite(
		demobEndpoint,
		suite,
		opts...,
	)

	return hellocli, err
}
