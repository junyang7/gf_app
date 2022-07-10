package _as

import (
	"strconv"
	"strings"
)

func Int64(data interface{}) int64 {
	switch data.(type) {
	case []byte:
		i64, err := strconv.ParseInt(strings.TrimSpace(string(data.(interface{}).([]byte))), 10, 64)
		if nil != err {
			return 0
		}
		return i64
	case string:
		i64, err := strconv.ParseInt(strings.TrimSpace(data.(interface{}).(string)), 10, 64)
		if nil != err {
			return 0
		}
		return i64
	case int8:
		return int64(data.(interface{}).(int8))
	case int16:
		return int64(data.(interface{}).(int16))
	case int32:
		return int64(data.(interface{}).(int32))
	case int64:
		return data.(interface{}).(int64)
	case int:
		return int64(data.(interface{}).(int))
	case uint8:
		return int64(data.(interface{}).(uint8))
	case uint16:
		return int64(data.(interface{}).(uint16))
	case uint32:
		return int64(data.(interface{}).(uint32))
	case uint64:
		return int64(data.(interface{}).(uint64))
	case uint:
		return int64(data.(interface{}).(uint))
	case float32:
		return int64(data.(interface{}).(float32))
	case float64:
		return int64(data.(interface{}).(float64))
	case bool:
		if data.(interface{}).(bool) {
			return 1
		}
		return 0
	default:
		return 0
	}
}
