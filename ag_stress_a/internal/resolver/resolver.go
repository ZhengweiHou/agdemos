package resolver

import (
	"ag-core/ag/ag_conf"
	"errors"
	"log/slog"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

var (
	FxNoProviderError = errors.New("no provider")
)

// BuildKitexHzwResolver 构建服务发现解析器
func BuildKitexHzwResolver(env ag_conf.IConfigurableEnvironment, namingClient naming_client.INamingClient) (discovery.Resolver, error) {

	if namingClient == nil {
		slog.Warn("Naming client is nil, resolver will not be created")
		return nil, FxNoProviderError
	}

	rtype := env.GetProperty("kitex.client.resolver.type")

	if rtype == "hzw" {
		return resolver.NewNacosResolver(namingClient,
			resolver.WithGroup("DEFAULT_GROUP"),
			resolver.WithCluster("DEFAULT")), nil

	}
	return nil, FxNoProviderError
}

// func HzwResolverConditional(env ag_conf.IConfigurableEnvironment) fx.Option {
// 	if env.GetProperty("kitex.client.resolver.type") == "hzw" {
// 		return fx.Provide(BuildKitexHzwResolver)
// 	}
// 	return fx.Options() // 空选项，相当于不加载
// }


