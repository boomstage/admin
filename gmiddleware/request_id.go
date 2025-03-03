package gmiddleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
)

const (
	// RequestIDHeader 请求ID头
	RequestIDHeader = "X-Request-Id"
)

/*
	需要用到request id的中间件，必须要保证此中间件先被注册
	例如:
		h.Use(gmiddleware.NewRequestID()) // request id在accesslog先注册
		h.Use(gmiddleware.NewAccesslog())
*/

// NewRequestID 生成请求ID
func NewRequestID() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取requestID, 发现没有自动填充
		requestID := string(c.Request.Header.Peek(RequestIDHeader))
		if len(requestID) <= 0 {
			requestID = uuid.New().String()
			c.Request.Header.Set(RequestIDHeader, requestID)
		}
		// 在上下文设置requestID，方便其他中间件使用
		c.Set(RequestIDHeader, requestID)
		// response也设置上,方便客户端跟踪
		c.Response.Header.Set(RequestIDHeader, requestID)
		// 设置到外部context中,方便链路传递
		ctx = context.WithValue(ctx, RequestIDHeader, requestID)
		// next
		c.Next(ctx)
	}
}
