package service

import "github.com/boomstage/admin/biz/model"

type UserSvc struct {
}

func InitUser() *UserSvc {
	return &UserSvc{}
}

func (u *UserSvc) CreateOrGetID(user *model.User) (int64, error) {
	return 1, nil
}
