package util

import (
	"fmt"
	"strconv"
)

func Val2string(val interface{}) string {
	switch val.(type) {
	case string:
		return val.(string)
	default:
		return fmt.Sprintf("%v", val)
	}
}

func Val2int(val interface{}) int {
	switch val.(type) {
	case int:
		return val.(int)
	case int8:
		return int(val.(int8))
	case int16:
		return int(val.(int16))
	case int32:
		return int(val.(int32))
	case int64:
		return int(val.(int64))
	case string:
		v1 := val.(string)
		v2, err := strconv.Atoi(v1)
		if err != nil {
			return 0
		}
		return v2
	case float32:
		v1 := val.(float32)
		return int(v1)
	case float64:
		v1 := val.(float64)
		return int(v1)
	default:
		return 0
	}
}

func Val2int64(val interface{}) int64 {
	return int64(Val2int(val))
}
