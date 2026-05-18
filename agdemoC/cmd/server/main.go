package main

import (
	"ag-core/ag/ag_app"
	"ag-core/ag/ag_service"
	"ag-core/contribute/agdb"
	"ag-core/contribute/agdb/gormdb"
	hertzServer "ag-core/contribute/aghertz/server"
	kitexServer "ag-core/contribute/agkitex/server"
	"ag-core/contribute/agnacos"
	"ag-core/contribute/agredis"
	"agdemoC/internal"
	"net/http"
	"os"

	"ag-core/fxs"

	"go.uber.org/fx"
	// _ "agdemoA/api/logs/logserver"
	_ "net/http/pprof"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string

	id, _ = os.Hostname()
)

func main() {

	// 启动 http pprof
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

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
	// agnacos.FxNacosConfigMode,
	agnacos.FxNacosNamingMode,
	agnacos.FxNacosConfigMode,

	// nacos远程配置
	agnacos.FxEnableNacosRemoteConfigMode,

	// 日志
	fxs.FxAgSlogZapMode,
	fxs.FxAgSlogMode,

	// 根APP
	fxs.FxAppMode,

	/** BaseServer **/
	hertzServer.FxAgHertzServerModule,
	kitexServer.FxKitexServerBaseModule,

	// 服务层
	ag_service.FxAgServiceMode,

	// 数据库
	agdb.FxAgDbModule,
	gormdb.FxAicGromdbModule,

	// redis
	agredis.FxAgRedisServerMode,

	/** 项目内部组件 **/

	// 项目内部组件(自定义)
	internal.FxInternalModule,
)
