//go:build js && wasm

package main

//go:generate cp $GOROOT/misc/wasm/wasm_exec.js .

import (
	"encoding/base64"
	"syscall/js"

	"github.com/skip2/go-qrcode"
)

func generateQRCode(this js.Value, p []js.Value) interface{} {
	input := p[0].String()
	qr, err := qrcode.New(input, qrcode.Medium)
	if err != nil {
		return nil
	}

	// Generate the QR code image as a PNG
	qrImage, err := qr.PNG(256)
	if err != nil {
		return nil
	}

	qrCodeBase64 := base64.StdEncoding.EncodeToString(qrImage)
	return js.ValueOf(qrCodeBase64)
}

func main() {
	ch := make(chan struct{})
	js.Global().Set("generateQRCode", js.FuncOf(generateQRCode))
	<-ch
}
