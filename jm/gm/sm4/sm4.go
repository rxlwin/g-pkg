package sm4

import (
	"bytes"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"github.com/rxlwin/g-pkg/jm/sm/md5"
	OriginSm4 "github.com/tjfoc/gmsm/sm4"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

// 最原始的加密
func JiaMi(key string, originByte []byte) ([]byte, error) {
	JiaMiByte, err1 := OriginSm4.Sm4Ecb(FormatKey(key), originByte, true)
	if err1 != nil {
		return nil, errors.New("加密失败:" + err1.Error())
	}
	return JiaMiByte, nil
}

// 最原始的解密
func JieMi(key string, JiaMiByte []byte) ([]byte, error) {
	JieMiByte, err := OriginSm4.Sm4Ecb(FormatKey(key), JiaMiByte, false)
	if err != nil {
		return nil, errors.New("解密错误：" + err.Error())
	}
	return JieMiByte, nil

}

// Base32加密
func JiaMiToBase32(key string, originByte []byte) (string, error) {
	JiaMiByte, err0 := JiaMi(key, originByte)
	if err0 != nil {
		return "", err0
	}
	return base32.HexEncoding.WithPadding(base32.NoPadding).EncodeToString(JiaMiByte), nil
}

// Base32解密
func Base32JieMi(key, JiaMiStr string) ([]byte, error) {
	JiaMiByte, err0 := base32.HexEncoding.WithPadding(base32.NoPadding).DecodeString(JiaMiStr)
	if err0 != nil {
		return nil, err0
	}
	return JieMi(key, JiaMiByte)
}

// 市特殊加密
func sJiaMi(key, originStr string) (string, error) {
	oByte := []byte(originStr)
	gbkByte, err0 := Utf8ToGbk(oByte)
	if err0 != nil {
		return "", errors.New("utf8转Gbk失败:" + err0.Error())
	}
	JiaMiByte, err1 := OriginSm4.Sm4Ecb([]byte(key), gbkByte, true)
	if err1 != nil {
		return "", errors.New("加密失败:" + err1.Error())
	}
	return base64.StdEncoding.EncodeToString(JiaMiByte), nil
}

// 市特殊解密
func sJieMi(key, JiaMiStr string) (string, error) {
	oJieMiByte, err := base64.StdEncoding.DecodeString(JiaMiStr)
	if err != nil {
		return "", errors.New("16进制字符转字节错误: " + err.Error())
	}
	JieMiByte, err := OriginSm4.Sm4Ecb([]byte(key), oJieMiByte, false)
	if err != nil {
		return "", errors.New("解密错误：" + err.Error())
	}
	//fmt.Println("解密后(byte)：", JieMiByte)
	utf8JieMi, err := GbkToUtf8(JieMiByte)
	return string(utf8JieMi), err
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(utf8Byte []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(utf8Byte), simplifiedchinese.GBK.NewEncoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func FormatKey(key string) []byte {
	return md5.ToBytes(key)
}
