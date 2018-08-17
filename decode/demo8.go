package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/sha1"
	"log"
	"math/big"
)

func main() {
	sha1 := sha1.New()
	n := new(big.Int)
	d := new(big.Int)

	rsa_modulus := "a8b3b284af8eb50b387034a860f146c4919f318763cd6c5598c8ae4811a1e0abc4c7e0b082d693a5e7fced675cf4668512772c0cbc64a742c6c630f533c8cc72f62ae833c40bf25842e984bb78bdbf97c0107d55bdb662f5c4e0fab9845cb5148ef7392dd3aaff93ae1e6b667bb3d4247616d4f5ba10d4cfd226de88d39f16fb"
	rsa_d := "53339cfdb79fc8466a655c7316aca85c55fd8f6dd898fdaf119517ef4f52e8fd8e258df93fee180fa0e4ab29693cd83b152a553d4ac4d1812b8b9fa5af0e7f55fe7304df41570926f3311f15c4d65a732c483116ee3d3d2d0af3549ad9bf7cbfb78ad884f84d5beb04724dc7369b31def37d0cf539e9cfcdd3de653729ead5d1"

	n.SetString(rsa_modulus, 16)
	d.SetString(rsa_d, 16)
	public := rsa.PublicKey{n, 65537}
	d.SetString(rsa_d, 16)
	private := new(rsa.PrivateKey)
	private.PublicKey = public
	private.D = d

	seed := []byte{0x18, 0xb7, 0x76, 0xea, 0x21, 0x06, 0x9d, 0x69,
		0x77, 0x6a, 0x33, 0xe9, 0x6b, 0xad, 0x48, 0xe1, 0xdd,
		0xa0, 0xa5, 0xef,
	}
	randomSource := bytes.NewReader(seed)

	in := []byte("Hello World")

	encrypted, err := rsa.EncryptOAEP(sha1, randomSource, &public, in, nil)
	if err != nil {
		log.Println("error: %s", err)
	}

	plain, err := rsa.DecryptOAEP(sha1, nil, private, encrypted, nil)
	if err != nil {
		log.Println("error: %s", err)
	}

	log.Println(string(plain))
}
