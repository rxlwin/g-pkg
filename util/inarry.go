package util

import (
	"reflect"
)

func InArray(val interface{}, list interface{}) bool {
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(list)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				//这里如果需要位置, 也可以返回 i
				return true
			}
		}
	default:
		return false
	}
	return false
}
