package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/boomstage/admin/biz/model"
	"github.com/boomstage/admin/biz/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/duke-git/lancet/v2/convertor"
	"math/rand"
)

var Google GoogleHandler

type GoogleHandler struct{}

// 生成随机状态参数（防止 CSRF 攻击）
func generateState() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// 处理 Google 登录请求
func (g *GoogleHandler) HandleGoogleLogin(ctx context.Context, c *app.RequestContext) {
	state := generateState()
	c.SetCookie("oauthstate", state, 600, "/", string(c.Host()), protocol.CookieSameSiteLaxMode, true, true) // 10 分钟有效
	url := model.GoogleOAuthConfig.AuthCodeURL(state)
	c.Redirect(302, []byte(url))
}

// 处理 Google 登录回调
func (g *GoogleHandler) HandleGoogleCallback(ctx context.Context, c *app.RequestContext) {
	// 检查状态参数
	state, _ := c.GetQuery("state")
	if state != string(c.Cookie("oauthstate")) {
		util.Zerolog.Error().Msg("Invalid state parameter")
		return
	}
	// 使用授权码换取 Access Token
	code, _ := c.GetQuery("code")
	token, err := model.GoogleOAuthConfig.Exchange(context.Background(), convertor.ToString(code))
	if err != nil {
		util.Zerolog.Err(err).Msg("Failed to exchange token")
		return
	}

	// 获取用户信息
	client := model.GoogleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		util.Zerolog.Err(err).Msg("Failed to get user info")
		return
	}
	defer resp.Body.Close()

	var userInfo struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	json.NewDecoder(resp.Body).Decode(&userInfo)
	util.FmtDataResp(c, userInfo)
	return
}
