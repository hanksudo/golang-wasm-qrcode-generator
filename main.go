package main

import (
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"log"
)

func main() {
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("qrcode.png", png, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

