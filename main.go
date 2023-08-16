package main

import (
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"log"
)

func main() {
	var png []byte
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("qrcode.png", png, 0644)
}

