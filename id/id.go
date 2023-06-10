package id

import (
	"math"
	"strings"
	"sync"
	"time"
)

const (
	// BaseTimeMs 取 2023-01-01 00:00:00 为基准时间 (时间戳: 1672502400000)
	baseTimeMs int64 = 1672502400000
)

var mux sync.Mutex
var incId int64

func GetId() int64 {
	var newId int64 = 0
	curMs := getCurTimeMs()
	baseMs := getBaseTimeMs()
	timeSeg := curMs - baseMs //取最后40bit
	newId += timeSeg
	incId := getIncrementId()
	newId = newId*10000 + incId
	return newId
}

func getCurTimeMs() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func getBaseTimeMs() int64 {
	return baseTimeMs
}

func getIncrementId() int64 {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	incId++
	if incId >= 10000 {
		incId = 0
	}
	return incId
}

var dic = "Y1FC2JDK9P3X4NST5AE6VR7M8BQW0GZH"

func Int64ToStr(num int64) string {
	new_num_str := ""
	var remainder int
	var remainder_string string
	for num != 0 {
		remainder = int((num % 32))
		remainder_string = string(dic[remainder])
		//if 76 > remainder && remainder > 9 {
		//	remainder_string = string(dic[remainder])
		//} else {
		//	remainder_string = strconv.Itoa(remainder)
		//}
		new_num_str = remainder_string + new_num_str
		num = num / 32
	}
	return new_num_str
}

func StrToInt64(str string) int64 {
	var new_num float64
	new_num = 0.0
	nNum := len(strings.Split(str, "")) - 1
	//for _, value := range strings.Split(str, "") {
	for _, value := range str {
		tmp := float64(findDicKey(string(value)))
		if tmp != -1 {
			new_num = new_num + tmp*math.Pow(float64(32), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int64(new_num)
}

func findDicKey(in string) int {
	result := -1
	for k, v := range dic {
		if in == string(v) {
			return k
			result = k
			break
		}
	}
	return result
}
