package qr

import (
	qrcode "github.com/skip2/go-qrcode"
)

func Generate(text string) ([]byte, error) {
	err := qrcode.WriteFile("http://baidu.com", qrcode.Medium, 256, "qr.png")
	if err != nil {
		return nil, err
	}

	return nil, nil
}
