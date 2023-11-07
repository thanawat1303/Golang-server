package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func EncryptBase64(textEncryt string) string {
	return base64.StdEncoding.EncodeToString([]byte(textEncryt))
}

func EncryptSha256(textEncryt string) string {
	hash := sha256.New()           //create sha256
	hash.Write([]byte(textEncryt)) //encrpt 256
	hashed := hash.Sum(nil)        //input

	return
}

func EncrytAES(key []byte, textEncryt string) string {
	block, err := aes.NewCipher(key)

	out := make([]byte, len(textEncryt))

	block.Encrypt(out, []byte(textEncryt))
	print(out)

	return En
}

func main() {
	aes.NewCipher()
	cipher.NewGCM()
	hmac.New()
	base64.NewEncoding()
	fmt.Print()
}
