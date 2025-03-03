package gmiddleware

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
)

func GetOrigins(env string) []string {
	if env == "prod" {
		return []string{
			"https://admin.gosh.com",
			"https://www.gosh.com",
			"https://h5.gosh.com",
			"https://gosh.com",

			// app的离线包方案用了这个，先暂时加上
			"http://127.0.0.1:8080",
		}
	}

	// 测试服
	return []string{
		"https://admin-test.gosh0.com",
		"https://admin-local.gosh0.com",
		"https://live-test.gosh0.com",
		"https://live-dev.gosh0.com",
		"https://live-local.gosh0.com",
		"https://h5-test.gosh0.com",
		"https://h5-local.gosh0.com",
		"https://gosh0.com",
		"https://www.gosh0.com",

		// app的离线包方案用了这个，先暂时加上
		"http://127.0.0.1:8080",
	}
}

func NewCors(env string) app.HandlerFunc {
	// origins := GetOrigins(env)
	// 跨域
	// config := cors.DefaultConfig()
	// config.AllowOrigins = origins
	// config.AllowCredentials = true
	// config.AllowWildcard = true
	// config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	// config.AllowHeaders = []string{"Content-Type, Authorization"}
	// config.MaxAge = 12 * time.Hour

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowWildcard = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	config.MaxAge = 12 * time.Hour
	return cors.New(config)
}
