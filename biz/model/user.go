package model

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"gorm.io/gorm"
)

// User 结构体
type User struct {
	ID        int64   `gorm:"primaryKey" json:"id"`
	Username  string  `gorm:"size:50;not null" json:"username" vd:"len($)>8"`
	Phone     string  `gorm:"size:50;unique;not null" json:"phone" vd:"len($)>8"`
	Password  string  `gorm:"size:255;not null" json:"password" vd:"len($)>8"`
	Salt      string  `gorm:"size:50;not null" json:"salt"`
	Role      string  `gorm:"type:enum('super_admin', 'manager', 'operator');default:'operator';not null" json:"role"`
	Rate      float64 `gorm:"type:decimal(10,4);default:0;not null" json:"rate"`
	ParentID  uint    `json:"parent_id"`
	CreatedAt int64   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64   `gorm:"autoUpdateTime" json:"updated_at"`
}

// EncryptPassword 使用 MD5 + Salt 进行密码加密
func EncryptPassword(password, salt string) string {
	hash := md5.New()
	hash.Write([]byte(password + salt))
	return hex.EncodeToString(hash.Sum(nil))
}

// HashPassword 生成密码（MD5 + Salt）
func (u *User) HashPassword() {
	u.Password = EncryptPassword(u.Password, u.Salt)
}

// CheckPassword 校验密码
func (u *User) CheckPassword(password string) bool {
	s := EncryptPassword(password, u.Salt)
	return u.Password == s
}

// CreateUser 创建用户
func (u *User) CreateUser(db *gorm.DB, user *User) error {
	user.HashPassword()
	return db.Create(user).Error
}

// GetUserByID 根据 ID 获取用户
func (u *User) GetUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByName 获取用户
func (u *User) GetUserByName(db *gorm.DB, username string) error {
	err := db.Where("username = ?", username).First(u).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser 更新用户信息
func (u *User) UpdateUser(db *gorm.DB, id uint, updates map[string]interface{}) error {
	// 不能直接更新密码，密码必须经过加密
	if _, ok := updates["password"]; ok {
		return errors.New("password update not allowed via this method")
	}

	return db.Model(&User{}).Where("id = ?", id).Updates(updates).Error
}

// DeleteUser 删除用户
func (u *User) DeleteUser(db *gorm.DB, id uint) error {
	return db.Delete(&User{}, id).Error
}

// ListUsers 获取用户列表（分页）
func (u *User) ListUsers(db *gorm.DB, page, size int) ([]User, error) {
	var users []User
	offset := (page - 1) * size
	err := db.Offset(offset).Limit(size).Find(&users).Error
	return users, err
}
