package gmiddleware

import (
	"context"
	"fmt"
	"github.com/boomstage/admin/biz/model"
	"github.com/boomstage/admin/biz/util"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v5"
)

var (
	_secrets      = map[model.UserSource]string{}
	_excluedPaths = map[string]bool{}
)

type AuthOptions struct {
	ExcluedPaths []string
}

type AuthOption func(*AuthOptions)

const (
	PathTypeApp      string = "app"
	PathTypeAdmin    string = "admin"
	PathTypeInternal string = "internal"

	AuthKey   = "Authorization"
	CookieKey = "token"
)

// WithExcluedPaths 设置排除的路由,不用鉴权,不支持*
func WithExcluedPaths(paths []string) AuthOption {
	return func(o *AuthOptions) {
		o.ExcluedPaths = paths
	}
}

// NewAuth 鉴权
// tokens 传入,不写死代码
func NewAuth(secrets map[model.UserSource]string, opts ...AuthOption) app.HandlerFunc {
	_secrets = secrets
	options := &AuthOptions{}
	for _, opt := range opts {
		opt(options)
	}

	if options.ExcluedPaths != nil {
		for _, path := range options.ExcluedPaths {
			_excluedPaths[path] = true
		}
	}

	return func(_ context.Context, c *app.RequestContext) {

		method := string(c.Method())
		if method == consts.MethodOptions {
			return
		}

		path := string(c.Path())
		// 排除的路由,不用鉴权
		if len(_excluedPaths) > 0 {
			if _, ok := _excluedPaths[path]; ok {
				return
			}
		}

		paths := strings.Split(path, "/")
		if len(paths) < 3 {
			log.Printf("path: %s can not access", path)
			c.AbortWithStatus(consts.StatusNotFound)
			return
		}
		pathType := paths[2]
		userSource := model.UserSource(util.ToInt64(c.Query("us")))

		switch pathType {
		case PathTypeInternal:
			ip := c.ClientIP()
			if !util.IsInternalIP(ip) {
				log.Printf("path: %s is internal path, uid: %s, ip: %s can not access", path, c.GetString("uid"), ip)
				c.AbortWithStatus(consts.StatusForbidden)
				return
			}
			// 如果是内网, 不需要鉴权, 但需要us为UserSourceSvc
			if userSource != model.UserSourceSvc {
				log.Printf("path: %s is internal path, uid: %s, us: %d can not access", path, c.GetString("uid"), int64(userSource))
				c.AbortWithStatus(consts.StatusForbidden)
				return
			}
			return
		case PathTypeAdmin:
			if userSource != model.UserSourceAdmin {
				log.Printf("path: %s is admin path, uid: %s, us: %d can not access", path, c.GetString("uid"), int64(userSource))
				c.AbortWithStatus(consts.StatusForbidden)
				return
			}
		}

		// 鉴权
		tokenString := string(c.GetHeader(AuthKey))
		if tokenString == "" {
			// get from cookie
			tokenString = string(c.Cookie(CookieKey))
		}
		token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(_secrets[userSource]), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(consts.StatusUnauthorized, &model.BaseResponse{
				Code:    consts.StatusUnauthorized,
				Message: "invalid token error",
			})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(consts.StatusUnauthorized, &model.BaseResponse{
				Code:    consts.StatusUnauthorized,
				Message: "invalid token",
			})
			return
		}

		claims, ok := token.Claims.(*AuthClaims)
		if !ok {
			c.AbortWithStatusJSON(consts.StatusUnauthorized, &model.BaseResponse{
				Code:    consts.StatusUnauthorized,
				Message: "invalid token, not ok",
			})
			return
		}

		// token 过期
		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			c.AbortWithStatusJSON(consts.StatusUnauthorized, &model.BaseResponse{
				Code:    consts.StatusUnauthorized,
				Message: "token expired",
			})
			return
		}

		uidStr, ok := c.GetQuery("uid")
		if !ok {
			c.AbortWithStatusJSON(consts.StatusUnauthorized, &model.BaseResponse{
				Code:    consts.StatusUnauthorized,
				Message: "query uid not found",
			})
			return
		}

		uid, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(consts.StatusUnauthorized, &model.BaseResponse{
				Code:    consts.StatusUnauthorized,
				Message: "parse uid int error",
			})
			return
		}
		if uid != claims.UID {
			c.AbortWithStatusJSON(consts.StatusUnauthorized, &model.BaseResponse{
				Code:    consts.StatusUnauthorized,
				Message: "query uid not match token uid",
			})
			return
		}

		// 附加上下文信息
		c.Set("gosh-uid", uid)
		c.Set("gosh-usource", userSource)
	}
}

func GenAuthToken(userSource model.UserSource, id int64, duration time.Duration) (token string, err error) {
	data := AuthClaims{
		id,
		userSource,
		jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration))},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	secret, ok := _secrets[userSource]
	if !ok {
		return "", fmt.Errorf("not found secret for %d", userSource)
	}
	token, err = t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// deprecated 因为前端要做SSR，在Next.js做SetCookie
func GenAuthTokenAndSetCookie(c *app.RequestContext, userSource model.UserSource, id int64, duration time.Duration) (token string, err error) {
	token, err = GenAuthToken(userSource, id, duration)
	if err != nil {
		return "", err
	}
	c.SetCookie(AuthKey, token, int(duration.Seconds()), "/", string(c.Host()), protocol.CookieSameSiteLaxMode, true, true)
	return token, nil

}

type AuthClaims struct {
	UID        int64            `json:"uid"`
	UserSource model.UserSource `json:"user_source"`
	jwt.RegisteredClaims
}
