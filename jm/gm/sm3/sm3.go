/*
对标Md5，比md5更安全，效率不如md5高
md5输出的是128byte（32字节）
sm3输出的是256byte（64字节）
*/

package sm3

import (
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
)

func To64(data string) string {
	s := sm3.New()
	s.Write([]byte(data))
	jmByte := s.Sum(nil)
	//fmt.Println(byteToString(jmByte))
	return hex.EncodeToString(jmByte)
}

func To32(data string) string {
	return To64(data)[0:32]
}

func To16(data string) string {
	return To64(data)[8:24]
}

// 与hex.EncodeToString效果一致
func byteToString(b []byte) string {
	ret := ""
	for i := 0; i < len(b); i++ {
		ret += fmt.Sprintf("%02x", b[i])
	}
	fmt.Println("ret = ", ret)
	return ret
}
