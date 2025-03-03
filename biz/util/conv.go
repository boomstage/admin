package util

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"

	json "github.com/bytedance/sonic"

	"golang.org/x/exp/rand"
)

func ToString(i interface{}) string {
	var value string
	switch i.(type) {
	case int, int8, int16, int32, int64:
		value = strconv.FormatInt(reflect.ValueOf(i).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		value = strconv.FormatUint(reflect.ValueOf(i).Uint(), 10)
	case float32:
		value = strconv.FormatFloat(reflect.ValueOf(i).Float(), 'f', 4, 32)
	case float64:
		value = strconv.FormatFloat(reflect.ValueOf(i).Float(), 'f', 4, 64)
	case string:
		value = i.(string)
	case []byte:
		value = string(i.([]byte))
	default:
		value = fmt.Sprintf("%v", i)
	}
	return value
}

func ToBytes(i interface{}) []byte {
	return []byte(ToString(i))
}

func IsInternalIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.IsPrivate()
}

func UniquIDs(ids []int64) []int64 {
	set := map[int64]bool{}
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		if _, ok := set[id]; !ok {
			result = append(result, id)
			set[id] = true
		}
	}
	return result
}

func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// func GenNanoID() string {
// 	id, err := gonanoid.New()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return id
// }

func InterfaceToJSONString(t interface{}) string {
	raw, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(raw)
}

func GetHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func UniqueList(list []int64) []int64 {
	set := map[int64]bool{}
	result := []int64{}
	for _, id := range list {
		if _, ok := set[id]; !ok {
			result = append(result, id)
			set[id] = true
		}
	}
	return result
}

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandString(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func StructToMap(u interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	v := reflect.ValueOf(u)

	for i := 0; i < v.NumField(); i++ {
		m[v.Type().Field(i).Name] = v.Field(i).Interface()
	}
	return m
}

// VersionCompare 版本比较, 在没有err的情况下相等返回0，v1比v2大返回1，反之返回-1
func VersionCompare(version1, version2 string) (int, error) {
	version1s := strings.Split(version1, ".")
	version2s := strings.Split(version2, ".")
	if len(version1s) != 3 || len(version2s) != 3 {
		return 0, errors.New("wrong version")
	}

	for i := 0; i < 3; i++ {
		v1 := version1s[i]
		version1Number, err := strconv.Atoi(v1)
		if err != nil {
			return 0, err
		}

		v2 := version2s[i]
		version2Number, err := strconv.Atoi(v2)
		if err != nil {
			return 0, err
		}

		if version1Number > version2Number {
			return 1, nil
		}
		if version1Number < version2Number {
			return -1, nil
		}
	}
	return 0, nil
}

func PanicRecover() {
	if err := recover(); err != nil {
		// Zerolog.Panic().Any("err", err).Msg("panic recover")
		return
	}
}

func MD5(sth string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(sth))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 获取分页起点
func GetPageStart(page, pageSize int64) int64 {
	if page <= 0 {
		page = 1
	}
	return (page - 1) * pageSize
}

func StructToQuery(s interface{}) (string, error) {
	u := url.Values{}

	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	// 检查输入是否为结构体
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("expected a struct, got %s", v.Kind())
	}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// 获取 `query` 标签
		tag := field.Tag.Get("query")
		if tag == "" {
			// 如果没有 `query` 标签，跳过
			continue
		}

		// 根据字段类型格式化值
		var strValue string
		switch value.Kind() {
		case reflect.String:
			strValue = value.String()
		case reflect.Int32, reflect.Int, reflect.Int64:
			strValue = fmt.Sprintf("%d", value.Int())
		case reflect.Bool:
			strValue = fmt.Sprintf("%t", value.Bool())
		default:
			continue
		}

		// 设置查询参数
		u.Set(tag, strValue)
	}

	return u.Encode(), nil
}
