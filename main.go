package main

import (
	"context"
	"fmt"
	"github.com/boomstage/admin/biz/dao"
	"github.com/boomstage/admin/biz/util"
	"time"

	"github.com/boomstage/admin/biz/mw"
	"github.com/boomstage/admin/biz/service"
	"github.com/boomstage/admin/gmiddleware"
	"github.com/cloudwego/hertz/pkg/app/server"
	_ "go.uber.org/automaxprocs"
)

func main() {
	start := time.Now().UnixMilli()
	EnvInit()
	dao.Init()
	//gfile.Init(dao.Conf.File)
	util.InitLog()
	service.Init()

	addr := ":8895"
	port := dao.Conf.Port
	if port > 0 {
		addr = fmt.Sprintf(":%d", port)
	}
	h := server.Default(
		server.WithHostPorts(addr),
	) // accesslog
	h.Use(gmiddleware.NewAccesslog(gmiddleware.WithZeroLog(&util.Zerolog)))
	// 跨域
	h.Use(gmiddleware.NewCors(GetEnv()))
	// 校验 token
	mw.InitAuth(h)

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
