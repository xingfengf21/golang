package main

import (
	// "bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	// "crypto/sha1"
	"fmt"
	// "log"
	// "math/big"
	"os"
)

const PRIVATE_KEY string = `
-----BEGIN RSA PRIVATE KEY-----
MIICYAIBAAKBgQCkRWOiMmFNQfV1sep8RD5rwjzwpyYCXIoWHM/2FiK812ps57D4
4S05N1V7Wf8pxdnABEGhwXtuXbSEnnqu+T0H4njJQ/Sadfru+IDtzU754zbfRh1t
74YC6QAabQp2kR1Pd3Syg3x73KsqDJ9Bze/fzrO0E0n+WwNP5KSC2b0qVwIDAQAB
AoGAMxfiBXDmEsGRNyo/ZKdVQu8HRVK338QorvgHNuNrqpDV6UmnIcHn7wPm8pbr
ZnLDZ5XrVAR5/7nli9o+zpWDQMG14BwDx4mhLdf1Y13ftKN0CMEOBxUA/VrnXqy5
uEtWpxLK6gwQxEOHQ4CbQBe7KQqBVwQ6OcAPfFTBOiuUpwECRQDNQvYFAlaaa0tz
B6chsptfZoc0uMxZmaDnW45lclMaHCQRSAc4Ekp/C8w507Ync1wel8J0RXxYk016
EPi/7JTAyXPr7QI9AMzgh/A94v4hP8JQLYWya388l+J8j4iCrQbinBap9tKhS4vi
IbqmEQoPrs1cMXmmLi/7uHF9bdgcU53O0wJEcRpEld4D+nLE0E2XOmpAUwEb32E0
uOp+Tv2UKTIluG1pMlyOHcd2ZSsLDZIq/PIEAJd+/rKLIgOI3YochyTFkrVuusEC
PQCsGxpXGMXtEeVLyKrVG1Luaep1monA9XWAyTCyYUk1G8Yo1mfbMdIyj4YagY4W
USGPsXf57pIt723WrX0CRCNX/lCPTL94XD7aVPy5ZfvbsezNMCjjO9qK+svMy0eJ
Uplq1dYZ8mfDmSnjVUe5BTIg6nNmrOkqnH/tHwK+RJqV/U/s
-----END RSA PRIVATE KEY-----
`

const PUBLIC_KEY string = `
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAKRFY6IyYU1B9XWx6nxEPmvCPPCnJgJcihYcz/YWIrzXamznsPjhLTk3
VXtZ/ynF2cAEQaHBe25dtISeeq75PQfieMlD9Jp1+u74gO3NTvnjNt9GHW3vhgLp
ABptCnaRHU93dLKDfHvcqyoMn0HN79/Os7QTSf5bA0/kpILZvSpXAgMBAAE=
-----END RSA PUBLIC KEY-----
`

func main() {
	// sha1 := sha1.New()
	// n := new(big.Int)
	// d := new(big.Int)

	// rsa_modulus := "a8b3b284af8eb50b387034a860f146c4919f318763cd6c5598c8ae4811a1e0abc4c7e0b082d693a5e7fced675cf4668512772c0cbc64a742c6c630f533c8cc72f62ae833c40bf25842e984bb78bdbf97c0107d55bdb662f5c4e0fab9845cb5148ef7392dd3aaff93ae1e6b667bb3d4247616d4f5ba10d4cfd226de88d39f16fb"
	// rsa_d := "53339cfdb79fc8466a655c7316aca85c55fd8f6dd898fdaf119517ef4f52e8fd8e258df93fee180fa0e4ab29693cd83b152a553d4ac4d1812b8b9fa5af0e7f55fe7304df41570926f3311f15c4d65a732c483116ee3d3d2d0af3549ad9bf7cbfb78ad884f84d5beb04724dc7369b31def37d0cf539e9cfcdd3de653729ead5d1"

	// n.SetString(rsa_modulus, 16)
	// d.SetString(rsa_d, 16)
	// public := rsa.PublicKey{n, 65537}
	// d.SetString(rsa_d, 16)
	// private := new(rsa.PrivateKey)
	// private.PublicKey = public
	// private.D = d

	// // seed := []byte{0x18, 0xb7, 0x76, 0xea, 0x21, 0x06, 0x9d, 0x69,
	// // 	0x77, 0x6a, 0x33, 0xe9, 0x6b, 0xad, 0x48, 0xe1, 0xdd,
	// // 	0xa0, 0xa5, 0xef,
	// // }
	// // randomSource := bytes.NewReader(seed)

	// in := []byte("Hello World")

	// encrypted, err := rsa.EncryptOAEP(sha1, nil, &public, in, nil)
	// if err != nil {
	// 	log.Println("error: %s", err)
	// }
	// log.Println(encrypted)
	// plain, err := rsa.DecryptOAEP(sha1, nil, private, encrypted, nil)
	// if err != nil {
	// 	log.Println("error: %s", err)
	// }

	// log.Println(string(plain))

	msg := []byte("The secret message!")
	label := []byte("")
	md5hash := md5.New()

	var privatekey rsa.PrivateKey
	var publickey rsa.PublicKey

	encryptedmsg, err := rsa.EncryptOAEP(md5hash, rand.Reader, &publickey, msg, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(msg), encryptedmsg)
	fmt.Println()

	// DecryptOAEP
	decryptedmsg, err := rsa.DecryptOAEP(md5hash, rand.Reader, &privatekey, encryptedmsg, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("OAEP decrypted [%x] to \n[%s]\n", encryptedmsg, decryptedmsg)

}
