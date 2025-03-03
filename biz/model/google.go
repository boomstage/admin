package model

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// 配置 Google OAuth 2.0
var GoogleOAuthConfig = &oauth2.Config{
	ClientID:     "382560526809-pudqrlcgcpal4osob3a0a15ah62q4gee.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "http://localhost:8895/admin/google/callback",
	Scopes:       []string{"email", "profile"},
	Endpoint:     google.Endpoint,
}

var JwtSecret = []byte("your_jwt_secret_key") // JWT 秘钥（请放到环境变量中）
