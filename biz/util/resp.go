package util

import (
	"github.com/boomstage/admin/biz/model"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"time"
)

func FmtOKResp(c *app.RequestContext) {
	c.JSON(consts.StatusOK, &model.BaseResponse{
		Code:       model.CodeOK,
		ServerTime: time.Now().Unix(),
	})
}

func FmtDataResp(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, &model.BaseResponse{
		Code:       model.CodeOK,
		Data:       data,
		ServerTime: time.Now().Unix(),
	})
}

func fmtDataResp(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, &model.BaseResponse{
		Code:       model.CodeOK,
		Data:       data,
		ServerTime: time.Now().Unix(),
	})
}
