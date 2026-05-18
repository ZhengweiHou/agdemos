package clients

import (
	"ag_stress_b/internal/adpgen/kitex/cservice"

	"ag-core/contribute/agkitex/client"

	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func NewCGrpcClient(suite *client.KitexClientSuite) (cservice.Client, error) {
	demobEndpoint := "ag_stress_c_grpc"

	opts := []kclient.Option{}
	opts = append(opts,
		kclient.WithSuite(tracing.NewClientSuite()),
		kclient.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "ag_stress_b"}),
	)

	bnfshcli, err := cservice.NewClientWithSuite(
		demobEndpoint,
		suite,
		opts...,
	)
	return bnfshcli, err
}
