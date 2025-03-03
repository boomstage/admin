package main

import (
	"github.com/boomstage/admin/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedAdminRegister(r *server.Hertz) {
	admin := r.Group("/admin")
	//admin.GET("/ping", handler.Ping)

	admin.GET("/google/login", handler.Google.HandleGoogleLogin)
	admin.GET("/google/callback", handler.Google.HandleGoogleCallback)
}
