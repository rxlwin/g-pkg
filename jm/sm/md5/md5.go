package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func ToBytes(s string) []byte {
	h := md5.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func ToStr32(s string) string {
	return hex.EncodeToString(ToBytes(s))
}

/*
MD5加密后的位数有两种类型：16位与32位，默认使用32位(32个字节)。
16位实际上是从32位字符串中取中间的第9位到第24位的部分
https://baijiahao.baidu.com/s?id=1732430062474523157&wfr=spider&for=pc
https://blog.51cto.com/u_14201949/6024188
*/
func ToStr16(s string) string {
	return ToStr32(s)[8:24]
}
