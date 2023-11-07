package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"encoding/base64"
	"fmt"
)

func main() {
	aes.NewCipher()
	cipher.NewGCM()
	hmac.New()
	base64.NewEncoding()
	fmt.Print()
}
