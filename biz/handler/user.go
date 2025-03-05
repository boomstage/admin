package handler

import (
	"context"
	"github.com/boomstage/admin/biz/dao"
	"github.com/boomstage/admin/biz/model"
	"github.com/boomstage/admin/biz/mw"
	"github.com/boomstage/admin/biz/util"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

var User UserHandler

type UserHandler struct {
}

// 用户注册(只允许超管,主管注册账户)
func (u *UserHandler) Create(ctx context.Context, c *app.RequestContext) {
	var req model.User
	if err := c.BindAndValidate(&req); err != nil {
		util.FmtErrResp(c, model.CodeParamInvalid, err.Error())
		return
	}
	// TODO check username or phone is exist

	// 生成随机 salt
	salt := util.GenerateSalt(8) // 生成8位随机盐

	// 创建用户
	user := &model.User{
		Username:  req.Username,
		Phone:     req.Phone,
		Salt:      salt,
		Role:      req.Role,
		Password:  req.Password,
		UpdatedAt: time.Now().Unix(),
		CreatedAt: time.Now().Unix(),
	}
	err := user.CreateUser(dao.DBM, user)
	if err != nil {
		util.FmtErrResp(c, model.CodeServerError, err.Error())
		return
	}
	util.FmtOKResp(c)
	return
}

func (u *UserHandler) Login(ctx context.Context, c *app.RequestContext) {
	type loginReq struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	var req loginReq
	if err := c.BindAndValidate(&req); err != nil {
		util.FmtErrResp(c, model.CodeParamInvalid, err.Error())
		return
	}
	user := &model.User{}
	err := user.GetUserByName(dao.DBM, req.UserName)
	if err != nil {
		util.FmtErrResp(c, model.CodeParamInvalid, err.Error())
		return
	}

	if !user.CheckPassword(req.Password) {
		util.FmtErrResp(c, model.CodeParamInvalid, "password err")
		return
	}
	// 生成 JWT Token
	token, err := mw.GenAuthTokenAndSetCookie(c, model.UserSourceApp, user.ID, 24*time.Hour)
	if err != nil {
		util.FmtErrResp(c, model.CodeServerError, "gen token err")
		return
	}
	type resp struct {
		Token string `json:"token"`
	}
	data := &resp{Token: token}
	util.FmtDataResp(c, data)
	return
}
