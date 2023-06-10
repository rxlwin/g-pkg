package sign

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"
)

//var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

//GetUUID 获取随机字符串
func GetUUID() (string, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b), nil
}

func Signature(secretKeyName, appSecret string, sendParamEntity interface{}) string {
	str := getFieldString(sendParamEntity)
	if str == "" {
		return ""
	}
	stringA := fmt.Sprintf("%s&%s=%s", str, secretKeyName, appSecret)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(stringA))
	sign := hex.EncodeToString(md5Ctx.Sum(nil))
	return strings.ToUpper(sign)
}

//GetFieldString 获取结构体字段及值的拼接值
func getFieldString(sendParamEntity interface{}) string {
	m := reflect.TypeOf(sendParamEntity)
	v := reflect.ValueOf(sendParamEntity)
	var tagName string
	numField := m.NumField()
	w := make([]string, numField)
	numFieldCount := 0
	for i := 0; i < numField; i++ {
		fieldName := m.Field(i).Name
		tags := strings.Split(string(m.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[1]
		} else {
			tagName = m.Field(i).Name
		}
		if tagName == "xml" {
			continue
		}
		fieldValue := v.FieldByName(fieldName).Interface()

		if fieldValue != "" {
			if strings.Contains(tagName, "omitempty") {
				tagName = strings.Split(tagName, ",")[0]
			}
			s := fmt.Sprintf("%s=%v", tagName, fieldValue)
			w[numFieldCount] = s
			numFieldCount++
		}
	}
	if numFieldCount == 0 {
		return ""
	}
	w = w[:numFieldCount]
	sort.Strings(w)
	return strings.Join(w, "&")
}
