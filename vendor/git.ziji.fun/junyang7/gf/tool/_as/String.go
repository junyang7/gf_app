package _as

import (
	"strconv"
)

func String(data interface{}) string {
	switch data.(type) {
	case []byte:
		return string(data.(interface{}).([]byte))
	case string:
		return data.(interface{}).(string)
	case int8:
		return strconv.FormatInt(int64(data.(interface{}).(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(data.(interface{}).(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(data.(interface{}).(int32)), 10)
	case int64:
		return strconv.FormatInt(data.(interface{}).(int64), 10)
	case int:
		return strconv.FormatInt(int64(data.(interface{}).(int)), 10)
	case uint8:
		return strconv.FormatUint(uint64(data.(interface{}).(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(data.(interface{}).(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(data.(interface{}).(uint32)), 10)
	case uint64:
		return strconv.FormatUint(data.(interface{}).(uint64), 10)
	case uint:
		return strconv.FormatUint(uint64(data.(interface{}).(uint)), 10)
	case float32:
		return strconv.FormatFloat(float64(data.(interface{}).(float32)), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(data.(interface{}).(float64), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(data.(interface{}).(bool))
	default:
		return ""
	}
}
