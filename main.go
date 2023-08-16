package main

import (
	"syscall/js"

	"github.com/skip2/go-qrcode"
)

var jsQRCode js.Value

func GenerateQRCode(input string) ([]byte, error) {
	png, err := qrcode.Encode(input, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return png, nil
}

func main() {

}
