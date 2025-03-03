package util

import (
	"reflect"
	"strconv"
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
