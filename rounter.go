package main

import (
	"github.com/boomstage/admin/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedAdminRegister(r *server.Hertz) {
	app := r.Group("/app")
	//admin.GET("/ping", handler.Ping)

	app.GET("/google/login", handler.Google.HandleGoogleLogin)
	app.GET("/google/callback", handler.Google.HandleGoogleCallback)
}
