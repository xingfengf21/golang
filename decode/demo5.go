package main

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	// "math/big"
	"net/http"
	// "reflect"
)

const PRIVATE_KEY string = `
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDESDVYOL3TIeBJ0nn0wpMrnX2d2AnfRRAQfdsXPxeEZqPAApKo
HCxSpNvjzlk0daadRBUPM2cDYMd4qhcE0iKjztRC9FxiKWxeu9uPmuuDPQ4vVB3D
cg493WSau/8I5yDOZ6d9T/vriM4SXjfaCSpfr+ARSaSuAWe99IWXkF8Y/QIDAQAB
AoGBAJQq7hnqu2Dl8anIeMeLGg2SlYDR1KH2QGoyl3C078vCr1oClzLQEaVRTaxp
2G0BYwTUoBEZfGnQZMNxom7BMeOjIn570TS9rVxLTz3JhKLD0Krp8xoT4PS31g9u
H8yTomEmgcnkVgRLnbPhx80qLeJ5GpI7cn4QBor2iSwCxuMBAkEA1S3glBEu+sNe
l2O4xTdCS+Ms5IL5KbJUIb8aCzowCNCAvhqPHPXCk/91TYYoP2X9vcAOPGB5IoJx
5tELerWo2wJBAOu1ccjdVDkU1HcJgpZvzVFGQWMyHgjMeALLOCQa+wtxJeADkpjc
jXXKT5KJtdW3Ee4qjQnuIk/NpNL7klCt4QcCQQC0iYveGcddm0NqlouOicPdnOES
DVGSXHIfukYGKBsx+9JsLPXyordMu1XOd8VyG6AOqK3luBYegbaQsY01OM9tAkBF
wW3KFw4s0mhKTisyYCXwDo2dUKCo0/IUSZX0wXNYIIUn9Ya17FeE63l/EXgZplMN
G9SfgrTuzxofQko7zCmTAkEAn0moEkIxAUfRDYfYW1mAGnGYS3NjDKKu8gj0vejh
V29mv5626/xiHsVJt2KofTP5KolOyX98kqydHQ2BJnibMg==
-----END RSA PRIVATE KEY-----
`

func Decrypt(ciphertext []byte) ([]byte, error) {
	fmt.Println("Decrypt 001")
	privateKey := []byte(PRIVATE_KEY)
	block, _ := pem.Decode(privateKey)
	fmt.Println("Decrypt 002")
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	fmt.Println("Decrypt 003 priv:", priv)
	fmt.Println("Decrypt ciphertext:", ciphertext)
	dec, err := rsa.DecryptOAEP(
		sha1.New(), nil, priv, ciphertext, nil)
	if err != nil {
		fmt.Println("xxxxx", err)
		// panic(err)
	}
	fmt.Println("Decrypt 006")
	return dec, err
}

const BFHC = "BFHC"

func urlsafeBase64Decode(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}

func GetMD5Hash(s []byte) string {
	hasher := md5.New()
	hasher.Write(s)
	return hex.EncodeToString(hasher.Sum(nil))
}

func decode_hc(s string) (string, []byte, error) {
	dhc, err := urlsafeBase64Decode(s)
	if err != nil {
		return "", nil, err
	}
	fmt.Println("dhc:", dhc)
	out, err := Decrypt(dhc)
	if err != nil {
		return "", nil, err
	}
	return string(out[:32]), out[32:], nil
}

func decrypt(hc string, data []byte) ([]byte, error) {
	h, c, err := decode_hc(hc)
	if err != nil {
		return nil, err
	}
	data = xor_decode(data, c)
	m := GetMD5Hash(data)
	if m != h {
		return nil, fmt.Errorf("MD5 sum error.")
	}
	return data, nil
}

func xor_decode(input []byte, key []byte) []byte {
	for i := range input {
		input[i] = input[i] ^ key[i%len(key)]
	}
	return input
}

func decrypt_bytes(body []byte) ([]byte, error) {
	d := bytes.SplitN(body, []byte(","), 2)
	if len(d) != 2 {
		return nil, fmt.Errorf("Segment Length Error.")
	}
	return decrypt(string(d[0]), d[1])
}

func decrypt_resp(resp *http.Response) ([]byte, error) {
	// 从response中解密出byte的数据
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%d, %s", resp.StatusCode, string(body))
	}
	return decrypt_bytes(body)
}

func Sign(ciphertext []byte) ([]byte, error) {

	privateKey := []byte(PRIVATE_KEY)
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	s := sha1.Sum(ciphertext)
	dec, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, s[:])
	if err != nil {
		panic(err)
	}
	return dec, err

}

func main() {
	// var publickey rsa.PublicKey
	// msg := []byte("hi")
	// // label := []byte("")

	// seed := []byte{0x18, 0xb7, 0x76, 0xea, 0x21, 0x06, 0x9d, 0x69,
	// 	0x77, 0x6a, 0x33, 0xe9, 0x6b, 0xad, 0x48, 0xe1, 0xdd,
	// 	0xa0, 0xa5, 0xef,
	// }
	// randomSource := bytes.NewReader(seed)
	// n := new(big.Int)
	// public := rsa.PublicKey{n, 65537}
	// encryptedmsg, err := rsa.EncryptOAEP(sha1.New(), randomSource, &public, msg, nil)
	// fmt.Println(encryptedmsg, err)
	// var key rsa.PublicKey
	// fmt.Println("key", &key)
	// var MSG = []byte("hello go scure udp")
	// var RSA_LABEL = []byte("server")
	// resp, e := rsa.EncryptOAEP(sha1.New(), rand.Reader, &key, MSG, RSA_LABEL)
	// fmt.Println(resp, e)

	// input := []byte("git@github.com/mrkt")
	// inp := []byte("GkMQnId0R6+HsrK77SLdNMHCv7Q9FeOXB8V0oaWE/NyKBs3p1MSky81GH580TkbK3VjpQ6VqmnXltP6V7HSn/pi6GDxn4koGlOwuZby0apjM/W977tL3VajH6E7MPtmLTYaXG5VHbjgtLZBpJMWQxKWNig32dRe3vVntPGmjVLg=")

	// fmt.Println(inp)
	// uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	// fmt.Println("uEnc", uEnc, reflect.ValueOf(uEnc))

	// d, b, err := decode_hc(uEnc)
	// fmt.Println(d, b, err)
	// url, err := urlsafeBase64Decode(uEnc)
	// fmt.Println("url", url, err)
	// fmt.Println(string(url))

	// fmt.Println(GetMD5Hash(input))
	// fmt.Println(decrypt("hc", body))
	// r := decrypt_bytes(body)
	// fmt.Println(r)

	// str2 := "hello"
	// data2 := []byte(str2)
	// fmt.Println(data2)
	// str2 = string(data2[:])
	// fmt.Println(str2)

	// s := []byte("hello,world")
	// for _, v := range bytes.SplitN(s, []byte(","), 2) {
	// 	fmt.Print(string(v) + "   ") //什么都不输出
	// }
	// fmt.Println()

	// hexStr := "48656c6c6f20576f726c64"
	// data, err := hex.DecodeString(hexStr)
	// fmt.Println("data222", data, err)
	// fmt.Println(string(data))

	// h := md5.Sum([]byte(data))
	// hexStr2 := fmt.Sprintf("%x", h)
	// fmt.Println("hexStr2", hexStr2)
	// fmt.Println(hex.DecodeString(hexStr2))
	type args struct {
		ciphertext []byte
	}
	cypher, _ := hex.DecodeString("5b933f4ee80e0de0ed0feb6b3f0d6b1619eb94ec5d950015b66e227846e74b094b6b2c86d3e9d6737402c3368b81ef614b0f0508bc8230e96ae41fe088306623d108ed98081722929d420cd2c8fa77ed607128648ef29cc2c8292dec4b1171d92117d1289b03a4a1d82d5c537281c5af40acad3e6a38392e6f36b5326e3fd71e")

	fmt.Println("cypher", cypher[:])
	fmt.Println("cypher string:", string(cypher[:]))

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"ssss", args{cypher}, []byte("90cd7a0abef4eeb565b340f148b537964\xdd\x88\xd0\xaa,#\xb6d\xb2Kb\xa2f 9\x10Wo]6\x1ef\xdc\xe8\x88\xeao@\x18\x98\x05"), false},
	}
	for _, tt := range tests {
		c := tt.args.ciphertext
		fmt.Println("cccc", c)
		got, err := Decrypt(c)
		fmt.Println("got:", got)
		if (err != nil) != tt.wantErr {

			return
		}
		fmt.Println(got, err)

	}
	fmt.Println(tests)

	// type args struct {
	// 	ciphertext []byte
	// }
	// tests := []struct {
	// 	name    string
	// 	args    args
	// 	want    []byte
	// 	wantErr bool
	// }{
	// 	{"Signtest", args{[]byte("hello")}, []byte("5b933f4ee80e0de0ed0feb6b3f0d6b1619eb94ec5d950015b66e227846e74b094b6b2c86d3e9d6737402c3368b81ef614b0f0508bc8230e96ae41fe088306623d108ed98081722929d420cd2c8fa77ed607128648ef29cc2c8292dec4b1171d92117d1289b03a4a1d82d5c537281c5af40acad3e6a38392e6f36b5326e3fd71e"), false},
	// }

	// for _, tt := range tests {
	// 	fmt.Println("tt.args.ciphertext", tt.args.ciphertext)
	// 	fmt.Println("tt.args.ciphertext", string(tt.args.ciphertext))
	// 	got, err := Sign(tt.args.ciphertext)
	// 	if (err != nil) != tt.wantErr {
	// 		// t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
	// 		fmt.Println(err, tt.wantErr)
	// 		return
	// 	}
	// 	b64got := []byte(base64.StdEncoding.EncodeToString(got))
	// 	fmt.Println("b64got", b64got)
	// 	fmt.Println("b64got string:", string(b64got))
	// 	fmt.Println("want:", tt.want)
	// 	if !reflect.DeepEqual(b64got, tt.want) {
	// 		// t.Errorf("Sign() = %v, want %v", b64got, tt.want)
	// 		fmt.Println(666)
	// 		fmt.Println(b64got, tt.want)
	// 	}

	// }

}
