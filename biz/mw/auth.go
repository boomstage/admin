package mw

import (
	"github.com/boomstage/admin/biz/dao"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitAuth(h *server.Hertz) {
	excluedePaths := []string{
		"/admin/google/login",
		"/admin/google/callback",
		"/ping",
	}

	h.Use(NewAuth(dao.Conf.JWT.Secrets, WithExcluedPaths(excluedePaths)))
}
