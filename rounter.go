package main

import (
	"github.com/boomstage/admin/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedAdminRegister(r *server.Hertz) {
	app := r.Group("/app")
	app.GET("/ping", handler.Ping.Ping)
	app.GET("/ping/login", handler.Ping.Login)

	app.GET("/google/login", handler.Google.HandleGoogleLogin)
	app.GET("/google/callback", handler.Google.HandleGoogleCallback)

	user := app.Group("/user")
	user.POST("/create", handler.User.Create)
	user.POST("/login", handler.User.Login)

}
