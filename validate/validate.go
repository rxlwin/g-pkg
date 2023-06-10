package validate

import (
	"errors"
	"fmt"
	"github.com/rxlwin/g-pkg/util"
	"reflect"
	"strconv"
)

type Rule [5]interface{}
type Rules []Rule
type FP []interface{}

/*
rules := Rules{
	{&id, "id", nil, "Str", []int{1, 0}},
	{&page, "page", nil, "Int", []int{1, 0}},
	{&num, "num", nil, "Int", []int{1, 0}},
	{&tp, "type", nil, "enum", []interface{}{[]int{1, 2}}},
}
*/

type iRule struct {
	param        interface{}
	key          string //map对应的Key
	fun          string //参数需要验证的函数信息
	funParam     interface{}
	defaultValue interface{} //参数默认值，如果为nil时，要求必填
}

// Param 参数格式
type Param map[string]interface{}

func SetParam(rq Param, rules Rules) error {
	//rq := getAllParam(c)
	f := f{}
	for _, rule := range rules {
		oParam := rule[0] //outParam
		key := rule[1].(string)
		defaultVal := rule[2]
		funName := rule[3].(string)
		funParam := rule[4].(FP)

		oParamValue := reflect.ValueOf(oParam).Elem()
		oParamType := oParamValue.Type()
		arg, ok := rq[key]
		//fmt.Println(arg)
		if ok {
			argTypeOf := reflect.TypeOf(arg)
			if oParamType != argTypeOf {
				//fmt.Println("oParamTypeString",oParamType.String())
				newArg := formatVal(arg, oParamType.String())
				if newArg == nil {
					return errors.New(key + " 获取参数错误")
				}
				oParamValue.Set(reflect.ValueOf(newArg))
			} else {
				oParamValue.Set(reflect.ValueOf(arg))
			}
		} else {
			if defaultVal == nil {
				return errors.New(key + " 参数必填")
			} else {
				oParamValue.Set(reflect.ValueOf(defaultVal))
			}
		}

		fun := reflect.ValueOf(f).MethodByName(funName)
		if fun.IsValid() == false {
			return errors.New(key + " 规则参数错误 方法" + funName + "() 不存在")
		}
		funParamNum := fun.Type().NumIn()
		ft := fun.Type().In(0)
		ftStr := ft.Name()
		if ftStr != "" && ft != oParamType {
			return errors.New(key + " 规则参数错误 方法" + funName + "()中 " + key + "需" + ftStr + "类型")
		}
		in := []reflect.Value{oParamValue}
		if len(funParam)+1 != funParamNum {
			return errors.New(key + " 规则参数错误 fun:" + funName + " 需" + val2string(funParamNum-1) + "个参数")
		}
		for i, param := range funParam {
			rt := fun.Type().In(i + 1)
			if rt != reflect.TypeOf(param) {
				return errors.New(key + " 规则参数错误 fun:" + funName + " 参" + val2string(i+1) + " 需" + rt.Name() + "类型")
			}
			in = append(in, reflect.ValueOf(param))
		}
		res := fun.Call(in)[0]
		if !res.Bool() {
			return errors.New(key + " 校验失败 fun:" + funName + "()")
		}
	}
	return nil
}

// ====私有函数====

func formatVal(val interface{}, targetType string) interface{} {
	switch targetType {
	case "string":
		return util.Val2string(val)
	case "int":
		return util.Val2int(val)
	//case "[]map[string]interface {}":
	//	return val.([]map[string]interface{})
	case "int64":
		return int64(util.Val2int(val))
	default:
		return nil
	}
}

func val2string(val interface{}) string {
	switch val.(type) {
	case string:
		return val.(string)
	default:
		return fmt.Sprintf("%v", val)
	}
}

func val2int(val interface{}) int {
	switch val.(type) {
	case int:
		return val.(int)
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
