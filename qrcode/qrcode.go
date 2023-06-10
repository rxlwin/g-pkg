package qrcode

import (
	"encoding/base64"
	"github.com/skip2/go-qrcode"
)

func GetQrCodeBase64(con string) (string, error) {
	qrCodeByte, err := qrcode.Encode(con, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(qrCodeByte), nil
}
