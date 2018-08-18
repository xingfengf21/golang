package main

import (
	"strconv"
	"fmt"
)

func ConverToBin(n int) string{
	result := ""
	for ;n>0; n /= 2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {

	fmt.Println(ConverToBin(5),ConverToBin(12))

}
