package gmiddleware

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/rs/zerolog"
)

type accesslogOptions struct {
	log *zerolog.Logger
}

// AccesslogOption access option
type AccesslogOption func(*accesslogOptions)

// WithZeroLog set zerolog
func WithZeroLog(log *zerolog.Logger) AccesslogOption {
	return func(ao *accesslogOptions) {
		ao.log = log
	}
}

// NewAccesslog new accesslog
func NewAccesslog(ops ...AccesslogOption) app.HandlerFunc {
	options := &accesslogOptions{}
	for _, o := range ops {
		o(options)
	}
	if options.log != nil {
		return func(ctx context.Context, c *app.RequestContext) {
			start := time.Now()
			c.Next(ctx)
			end := time.Now()
			cost := end.Sub(start)

			var bizCode int
			bizCodeVal, isExist := c.Get("biz_code")
			if isExist {
				bizCode = bizCodeVal.(int)
			}

			statusCode := c.Response.StatusCode()

			var requestID string
			requestIDVal, isExist := c.Get(RequestIDHeader)
			if isExist {
				requestID = requestIDVal.(string)
			}

			uid := c.GetInt64("uid")
			seq, _ := c.GetQuery("seq")

			options.log.Info().
				Str("clt_ip", c.ClientIP()).
				Str("method", string(c.Method())).
				Str("path", string(c.Path())).
				Str("args", c.QueryArgs().String()).
				Str("seq", seq).
				Int64("uid", uid).
				Str("reqid", requestID).
				Int64("bcode", int64(bizCode)).
				Int64("scode", int64(statusCode)).
				Str("cost", fmt.Sprintf("%13v", cost)).
				Str("body", string(c.Request.Body())).
				Int("reslen", len(c.Response.Body())).
				Msg("access_log")
		}
	}

	// 没有设置log,按照原逻辑打印日志
	accesslog.Tags["biz_code"] = func(output accesslog.Buffer, c *app.RequestContext, _ *accesslog.Data, _ string) (int, error) {
		value, isExist := c.Get("biz_code")
		if !isExist {
			return output.WriteString("not_found")
		}

		return output.WriteString(strconv.Itoa(value.(int)))
	}

	accesslog.Tags["clt_ip"] = func(output accesslog.Buffer, c *app.RequestContext, _ *accesslog.Data, _ string) (int, error) {
		ip := c.ClientIP()
		return output.WriteString(ip)
	}

	accesslog.Tags["request_id"] = func(output accesslog.Buffer, c *app.RequestContext, data *accesslog.Data, extraParam string) (int, error) {
		value, isExist := c.Get(RequestIDHeader)
		if !isExist {
			return output.WriteString("not_found")
		}
		return output.WriteString(value.(string))
	}

	return accesslog.New(
		accesslog.WithFormat("access_log - ${clt_ip} - ${method} - ${path}?${queryParams} - body: ${body} - ${latency}, http_code: ${status}, biz_code: ${biz_code}, request_id: ${request_id}"),
	)
}
