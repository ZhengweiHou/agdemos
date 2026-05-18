package clients

import (
	"ag-core/contribute/aghertz/aghertzclient"
	"ag_stress_a/internal/adpgen/hertz/bservice"

	hclient "github.com/cloudwego/hertz/pkg/app/client"
)

func NewBHttpClient(cli *hclient.Client) bservice.BServiceHertzClient {
	demobEndpoint := "ag_stress_b_http"
	// demobEndpoint := "localhost:9922"

	bnfshcli := bservice.NewBServiceHertzClient(
		cli,
		aghertzclient.WithEndpoint(demobEndpoint),
		aghertzclient.WithSD(true), // 开启服务发现
	)
	return bnfshcli
}
