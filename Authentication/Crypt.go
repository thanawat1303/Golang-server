package main

import (
	"crypto/aes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
)

func EncryptBase64(textEncryt string) string {
	return base64.StdEncoding.EncodeToString([]byte(textEncryt))
}

func EncryptMd5(textEncryt string) string {
	hash := md5.New()
	hash.Write([]byte(textEncryt))
	hashed := hash.Sum(nil)

	return hex.EncodeToString(hashed)
}

func EncryptSha1(textEncryt string) string {
	hash := sha1.New()             //create sha256
	hash.Write([]byte(textEncryt)) //write encrpt sha256 to hash
	hashed := hash.Sum(nil)        //encrption sha256

	return hex.EncodeToString(hashed)
}

func EncryptSha256(textEncryt string) string {
	hash := sha256.New()
	hash.Write([]byte(textEncryt))
	hashed := hash.Sum(nil)

	return hex.EncodeToString(hashed)
}

func EncryptSha512(textEncryt string) string {
	hash := sha512.New()
	hash.Write([]byte(textEncryt))
	hashed := hash.Sum(nil)

	return hex.EncodeToString(hashed)
}

func EncrytAES(key []byte, textEncryt string) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err.Error()
	}

	out := make([]byte, len(textEncryt))

	block.Encrypt(out, []byte(textEncryt))

	return hex.EncodeToString(out)
}

func EncryptHmac(algorithm string, key []byte, textEncryt string) string {
	var algorithmIn func() hash.Hash

	switch algorithm {
	case "sha256":
		algorithmIn = sha256.New
	case "sha512":
		algorithmIn = sha512.New
	case "sha1":
		algorithmIn = sha1.New
	case "md5":
		algorithmIn = md5.New
	default:
		return "Inv algorithm"
	}

	encrpt := hmac.New(algorithmIn, key)
	encrpt.Write([]byte(textEncryt))

	return hex.EncodeToString(encrpt.Sum(nil))
}

func main() {
	fmt.Println(EncrytAES([]byte("thanawat1303"), "Hello"))
	// aes.NewCipher()
	// cipher.NewGCM()
	// hmac.New()
	// base64.NewEncoding()
	// fmt.Print()
}
