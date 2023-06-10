package res

const (
	SUCCESS = 200
)

var codeList = map[int]string{
	SUCCESS: "",
	0:       "网络故障",
	999:     "API Not Found",
}

func GetMsgByCode(code int) string {
	msg, ok := codeList[code]
	if !ok {
		msg = "未知错误"
	}
	return msg
}
