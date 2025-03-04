package service

import (
	"github.com/boomstage/admin/biz/dao"
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
	return token.SignedString(dao.Conf.JWT.Secrets[model.UserSourceGoogle])
}
