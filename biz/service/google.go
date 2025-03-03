package service

import (
	"github.com/boomstage/admin/biz/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type GoogleService struct {
}

func InitGoogle() *GoogleService {
	return &GoogleService{}
}

// CreateJWT 生成 JWT Token
func (s *GoogleService) CreateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // 24 小时有效期
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(model.JwtSecret)
}

// 保护的 API 路由
//func protectedHandler(w http.ResponseWriter, r *http.Request) {
//	tokenStr := r.Header.Get("Authorization")
//	if tokenStr == "" {
//		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
//		return
//	}
//
//	// 解析 JWT Token
//	claims := jwt.MapClaims{}
//	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//
//	if err != nil || !token.Valid {
//		http.Error(w, "Invalid token", http.StatusUnauthorized)
//		return
//	}
//
//	// 返回受保护的资源
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(map[string]string{
//		"message": "Access granted!",
//		"user":    claims["email"].(string),
//	})
//}

//func main() {
//	http.HandleFunc("/auth/google/login", handleGoogleLogin)
//	http.HandleFunc("/auth/google/callback", handleGoogleCallback)
//	http.HandleFunc("/protected", protectedHandler)
//
//	fmt.Println("Server started at http://localhost:8080")
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
