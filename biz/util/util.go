package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func ToInt64(i interface{}) int64 {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(v.Uint())
	case reflect.Float32, reflect.Float64:
		return int64(v.Float())
	case reflect.String:
		val, _ := strconv.ParseInt(v.String(), 10, 64)
		return val
	case reflect.Bool:
		if v.Bool() {
			return 1
		}
		return 0
	default:
		return 0
	}
}

// EncryptPassword 使用 MD5 + Salt 进行密码加密
func EncryptPassword(password, salt string) string {
	hash := md5.New()
	hash.Write([]byte(password + salt)) // 组合密码和 salt
	return hex.EncodeToString(hash.Sum(nil))
}

// 随机字符集
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateSalt 生成指定长度的随机 salt
func GenerateSalt(length int) string {
	rand.Seed(time.Now().UnixNano())
	salt := make([]byte, length)
	for i := range salt {
		salt[i] = charset[rand.Intn(len(charset))]
	}
	return string(salt)
}
