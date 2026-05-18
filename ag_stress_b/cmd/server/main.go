package main

import (
	"ag_stress_b/internal"

	"ag-core/ag/ag_app"
	"ag-core/ag/ag_log"

	"ag-core/fxs"

	"ag-core/ag/ag_service"
	hclient "ag-core/contribute/aghertz/client"
	hserver "ag-core/contribute/aghertz/server"
	kclient "ag-core/contribute/agkitex/client"
	kserver "ag-core/contribute/agkitex/server"
	"ag-core/contribute/agnacos"

	"go.uber.org/fx"

	// _ "ag_stress_b/api/logs/logserver"
	_ "go.uber.org/automaxprocs"
)

func main() {
	var fxopts []fx.Option

	fxopts = append(
		fxopts,
		mainFx,
		fx.Invoke(func(s *ag_app.App) {}),
	)

	fxapp := fx.New(
		fxopts...,
	)

	fxapp.Run()
}

var mainFx = fx.Module("main",

	/** conf **/
	// 初始化配置
	fxs.FxAgConfModule,

	// nacosconf
	agnacos.FxNacosConfigMode,
	agnacos.FxEnableNacosRemoteConfigMode,
	agnacos.FxNacosNamingMode,

	// 日志模块
	ag_log.FxAglogMode,

	// 根APP
	fxs.FxAppMode,

	// http服务
	hserver.FxAgHertzServerModule,
	// grpc服务
	kserver.FxKitexServerBaseModule,

	hclient.FxModuleAgHertzClient,
	kclient.FxKitexClientBaseModule,

	// 服务基础模块
	ag_service.FxAgServiceMode,

	// 数据库模块
	// agdb.FxAgDbModule,
	// gormdb.FxAicGromdbModule,

	// 内部组件
	internal.FxInternalModule,
)
