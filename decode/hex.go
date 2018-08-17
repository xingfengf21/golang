package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	str := []byte("Hello World")

	dstlen := hex.EncodedLen(len(str))
	dst := make([]byte, dstlen)

	outlen := hex.Encode(dst, str)
	fmt.Println("dstlen = ", dstlen, ", outlen = ", outlen)
	/*H 的 16 进制为  0x48 ->输出到 []byte 里面对于的为 52 和 58
	  4 对应 ASCII 码表中的10进制  52
	*/
	fmt.Println("hex = ", dst)

	hex2strlen := hex.DecodedLen(len(dst))
	str2 := make([]byte, hex2strlen)
	hex.Decode(str2, dst)
	fmt.Println("str2 = ", string(str2), ", dst = ", dst)

	estring := hex.EncodeToString([]byte(str))
	fmt.Println("estring = ", estring)
	fmt.Println("estring = ", []byte(estring))

	ebyte, err := hex.DecodeString(estring)
	if err != nil {
		fmt.Println("hex.DecodeString Error! ", err.Error())
	} else {
		fmt.Println("ebyte = ", string(ebyte))
	}

	fmt.Println(hex.Dump(ebyte))
}
