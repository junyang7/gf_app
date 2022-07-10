package _as

import (
	"strings"
)

func Bool(data interface{}) bool {
	switch data.(type) {
	case []byte:
		return len(data.(interface{}).([]byte)) > 0
	case string:
		return "" != strings.TrimSpace(data.(interface{}).(string))
	case int8:
		return data.(interface{}).(int8) > 0
	case int16:
		return data.(interface{}).(int16) > 0
	case int32:
		return data.(interface{}).(int32) > 0
	case int64:
		return data.(interface{}).(int64) > 0
	case int:
		return data.(interface{}).(int) > 0
	case uint8:
		return data.(interface{}).(uint8) > 0
	case uint16:
		return data.(interface{}).(uint16) > 0
	case uint32:
		return data.(interface{}).(uint32) > 0
	case uint64:
		return data.(interface{}).(uint64) > 0
	case uint:
		return data.(interface{}).(uint) > 0
	case float32:
		return data.(interface{}).(float32) > 0
	case float64:
		return data.(interface{}).(float64) > 0
	case bool:
		return data.(interface{}).(bool)
	default:
		return false
	}
}
