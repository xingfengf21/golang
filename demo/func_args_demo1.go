package main

import (
	"fmt"
	"math"
)

func myFunc(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}

}
func main() {

	myFunc(1, 2, 3, 4)
	myFunc(6, 7)
	fmt.Println(math.Max(1, 2))
}
