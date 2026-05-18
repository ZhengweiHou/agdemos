package clients

import (
	"ag-core/contribute/aghertz/aghertzclient"
	"ag_stress_b/internal/adpgen/hertz/cservice"

	hclient "github.com/cloudwego/hertz/pkg/app/client"
)

func NewCHttpClient(cli *hclient.Client) cservice.CServiceHertzClient {
	demobEndpoint := "ag_stress_c_http"
	// demobEndpoint := "localhost:9922"

	bnfshcli := cservice.NewCServiceHertzClient(
		cli,
		aghertzclient.WithEndpoint(demobEndpoint),
		aghertzclient.WithSD(true), // 开启服务发现
	)
	return bnfshcli
}
