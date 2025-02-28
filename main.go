package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hertz-contrib/pprof"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/goshgame/gcomponent/env"
	gfile "github.com/goshgame/gcomponent/file"
	"github.com/goshgame/gcomponent/general_conf"
	"github.com/goshgame/gcomponent/list"
	gtranslate "github.com/goshgame/gcomponent/translate"
	"github.com/goshgame/gmiddleware"
	"github.com/goshgame/gosh_admin/biz/dao"
	"github.com/goshgame/gosh_admin/biz/mw"
	"github.com/goshgame/gosh_admin/biz/service"
	"github.com/goshgame/gosh_admin/biz/util"
	_ "go.uber.org/automaxprocs"
)

func main() {
	start := time.Now().UnixMilli()
	env.Init()
	dao.Init()
	gfile.Init(dao.Conf.File)
	gtranslate.Init(dao.DBAdm, "server")
	util.InitLog()
	list.Init(dao.DBAdm)
	general_conf.Init(dao.DBAdm)
	service.Init()

	addr := ":8888"
	port := dao.Conf.Port
	if port > 0 {
		addr = fmt.Sprintf(":%d", port)
	}
	h := server.Default(
		server.WithHostPorts(addr),
	)
	// request id
	h.Use(gmiddleware.NewRequestID())
	// accesslog
	h.Use(gmiddleware.NewAccesslog(gmiddleware.WithZeroLog(&util.Zerolog)))
	// 跨域
	h.Use(gmiddleware.NewCors(env.GetEnv()))
	// 校验 token
	mw.InitAuth(h)
	// metrics
	h.Use(gmiddleware.NewServerMetric())
	gmiddleware.InitMetrics(dao.Conf.MetricPort)
	pprof.Register(h, "/gosh_admin/dev/pprof")

	// 记录操作日志
	mw.InitOperationLogMiddleware(h)

	// 设置监听关闭信号,优雅关闭
	h.SetCustomSignalWaiter(gmiddleware.WaitSignal)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		fmt.Println("Admin OnShutdown start")
		util.Zerolog.Info().Msg("Admin OnShutdown start")
		time.Sleep(3 * time.Second)
		util.Zerolog.Info().Msg("Admin OnShutdown end")
		fmt.Println("Admin OnShutdown end")
	})

	register(h)
	end := time.Now().UnixMilli()
	fmt.Println("*****start Admin server cost millis:", end-start)
	h.Spin()
}
