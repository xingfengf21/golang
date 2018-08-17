package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func urlsafeBase64Decode(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}

func GetMD5Hash(s []byte) string {
	hasher := md5.New()
	hasher.Write(s)
	fmt.Println(hasher.Sum(nil))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	input := []byte("hello golang base64 快乐编程http://www.01happy.com +~")
	uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	b, err := urlsafeBase64Decode(uEnc)
	fmt.Println(b)
	fmt.Println(err)

	bs := []byte("123")
	fmt.Println("bs", bs)
	fmt.Println(GetMD5Hash(bs))

}
