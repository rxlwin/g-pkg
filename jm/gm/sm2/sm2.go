package sm2

import (
	"encoding/hex"
	osm2 "github.com/ZZMarquis/gm/sm2"
)

type key struct {
	pubKey *osm2.PublicKey
	priKey *osm2.PrivateKey
}

func New(priKeyStr, pubKeyStr string) (*key, error) {
	pubKeyByte, err := hex.DecodeString(pubKeyStr)
	if err != nil {
		return nil, err
	}
	pubKeyByte = formatKey(pubKeyByte)
	pubKey, err := osm2.RawBytesToPublicKey(pubKeyByte)
	if err != nil {
		return nil, err
	}
	priKeyByte, err := hex.DecodeString(priKeyStr)
	if err != nil {
		return nil, err
	}
	priKeyByte = formatKey(priKeyByte)
	//t.Log(priKeyByte)
	priKey, err := osm2.RawBytesToPrivateKey(priKeyByte)
	if err != nil {
		return nil, err
	}
	return &key{
		pubKey: pubKey,
		priKey: priKey,
	}, nil
}

func (k *key) JiaMiC1C2C3(originStr string) (string, error) {
	jmByte, err := osm2.Encrypt(k.pubKey, []byte(originStr), osm2.C1C2C3)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(jmByte), nil
}

func (k *key) JieMiC1C2C3(JiaMiStr string) (string, error) {
	dataByte, err := hex.DecodeString(JiaMiStr)
	if err != nil {
		return "", err
	}
	plainText, err := osm2.Decrypt(k.priKey, dataByte, osm2.C1C2C3)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

/*
有时有些java版本的sm2生成的私钥，在某些情况下可能会出现首位0x00的情况。公钥首位可能出现0x04，
所以需要处理一下。
https://github.com/dromara/hutool/issues/2001
https://blog.csdn.net/weixin_49984575/article/details/125909844
https://www.cnblogs.com/goodAndyxublog/p/15654531.html
*/
func formatKey(keyByte []byte) []byte {
	switch len(keyByte) {
	case 33:
		if keyByte[0] == 0x00 {
			return keyByte[1:]
		}
	case 65:
		if keyByte[0] == 0x04 {
			return keyByte[1:]
		}
	}
	return keyByte
}
