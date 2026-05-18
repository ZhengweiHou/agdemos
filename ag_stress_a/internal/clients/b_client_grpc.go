package clients

import (
	"ag_stress_a/internal/adpgen/kitex/bservice"

	"ag-core/contribute/agkitex/client"

	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func NewBGrpcClient(suite *client.KitexClientSuite) (bservice.Client, error) {
	demobEndpoint := "ag_stress_b_grpc"

	opts := []kclient.Option{}
	opts = append(opts,
		kclient.WithSuite(tracing.NewClientSuite()),
		kclient.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "ag_stress_a"}),
	)

	bnfshcli, err := bservice.NewClientWithSuite(
		demobEndpoint,
		suite,
		opts...,
	)
	return bnfshcli, err
}
