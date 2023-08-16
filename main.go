//go:build js && wasm

package main

//go:generate cp $GOROOT/misc/wasm/wasm_exec.js .

import (
	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(input string) ([]byte, error) {
	png, err := qrcode.Encode(input, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return png, nil
}

func main() {

}
