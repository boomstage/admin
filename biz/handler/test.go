package handler

import (
	"context"
	"github.com/boomstage/admin/biz/util"
	"github.com/cloudwego/hertz/pkg/app"
)

var Test TestHandler

type TestHandler struct{}

func (p *TestHandler) Login(ctx context.Context, c *app.RequestContext) {
	util.FmtOKResp(c)
	return
}
