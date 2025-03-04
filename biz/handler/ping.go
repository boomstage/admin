package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

var Ping PingHandler

type PingHandler struct{}

func (p *PingHandler) Ping(ctx context.Context, c *app.RequestContext) {
	c.String(200, "pong")
}

func (p *PingHandler) Login(ctx context.Context, c *app.RequestContext) {
	c.String(200, "ok")
}
