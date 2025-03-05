package mw

import (
	"github.com/boomstage/admin/biz/dao"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitAuth(h *server.Hertz) {
	excluedePaths := []string{
		"/app/google/login",
		"/app/google/callback",
		"/ping",
		"/app/user/create",
		"/app/user/login",
	}

	h.Use(NewAuth(dao.Conf.JWT.Secrets, WithExcluedPaths(excluedePaths)))
}
