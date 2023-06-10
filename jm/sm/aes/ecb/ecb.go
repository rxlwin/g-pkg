/**
 * @author 张文兵
 * @mail wenbing@mgtv.com
 * @blog https://zhangwenbing.com/
 * @datetime 2020-08-06 19:40:49
 * @description
 * https://www.zhangwenbing.com/blog/golang/LV0Lvh_Ud
 */

package ecb

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

// Tool ECB工具的加密解密
// key字节长度16
type Tool struct {
	Key string
}

// NewTool 新建一个ECB加密工具
func NewTool(key string) *Tool {
	return &Tool{Key: key}
}

func (ecbTool *Tool) padding(src []byte) []byte {
	//填充个数
	paddingCount := aes.BlockSize - len(src)%aes.BlockSize
	if paddingCount == aes.BlockSize {
		return src
	}
	//填充数据
	return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
}

// unpadding
func (ecbTool *Tool) unPadding(src []byte) []byte {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
}

// Encrypt ECB加密
func (ecbTool *Tool) Encrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(ecbTool.Key))
	if err != nil {
		return nil, err
	}

	//padding
	src = ecbTool.padding(src)

	decrypted := make([]byte, len(src))
	size := aes.BlockSize

	for bs, be := 0, size; bs < len(src); bs, be = bs+size, be+size {
		if src[bs] == 0 {
			continue
		}
		block.Encrypt(decrypted[bs:be], src[bs:be])
	}

	return ecbTool.unPadding(decrypted), nil
}

// Decrypt ECB解密
func (ecbTool *Tool) Decrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(ecbTool.Key))
	if err != nil {
		return nil, err
	}

	//padding
	src = ecbTool.padding(src)

	decrypted := make([]byte, len(src))
	size := aes.BlockSize

	for bs, be := 0, size; bs < len(src); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], src[bs:be])
	}

	return ecbTool.unPadding(decrypted), nil
}

// Base64 base64加密
func (ecbTool *Tool) Base64(src []byte) string {
	return base64.RawURLEncoding.EncodeToString(src)
}

// UnBase64 base64解密
func (ecbTool *Tool) UnBase64(src string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(src)
}
