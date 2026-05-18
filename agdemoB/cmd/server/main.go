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
	"agdemoB/internal"
	"agdemoB/internal/adpgen"
	"agdemoB/internal/repository"
	"agdemoB/internal/svcgen"
	"flag"
	"fmt"
	"os"
	"runtime/pprof"

	"ag-core/fxs"

	"go.uber.org/fx"
	// _ "agdemoA/api/logs/logserver"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	threadProfile := pprof.Lookup("threadcreate")
	fmt.Printf(" beforeClient threads counts: %d\n", threadProfile.Count())
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

	fmt.Println("========shutdown======")
	fmt.Printf(" afterClient threads counts: %d\n", threadProfile.Count())
}

var mainFx = fx.Module("main",
	/** conf **/
	// 初始化配置
	fxs.FxAgConfModule,

	// nacosconf
	agnacos.FxNacosConfigMode,
	agnacos.FxNacosNamingMode,
	agnacos.FxEnableNacosRemoteConfigMode,

	// 根APP
	fxs.FxAppMode,

	fxs.FxAgSlogZapMode,

	fxs.FxAgSlogMode,

	ag_service.FxAgServiceMode,

	/** BaseServer **/
	hertzServer.FxAgHertzServerModule,
	kitexServer.FxKitexServerBaseModule,

	/** 其他组件 **/
	// 数据库配置
	agdb.FxAgDbModule,
	gormdb.FxAicGromdbModule,
	agredis.FxAgRedisServerMode,

	/** 项目内部组件 **/
	// 适配层
	adpgen.FxAdapterModule(),

	// 业务层实现
	svcgen.FxServiceWithProxyModule(), // 业务层实现,包含代理层
	// svcgen.FxServiceModule(), // 业务层实现，不包含代理层
	// svchzw.FxServiceWithProxyModule(), // hzw 测试版

	internal.FxInternalModule,

	repository.FxAgdemobRepositoryModule,
)
