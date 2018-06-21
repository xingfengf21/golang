package main

import (
	"fmt"
)

func main() {
	v1 := 2
	fmt.Println(v1)

	var v2 = 3
	fmt.Println(v2)
	// v3 := {1,2,3,4,5,6}
	// fmt.Println(v3)
	var balance = [10]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)

	// var v3 [10]int

	// for i := 0; i < 10; i++ {
	// 	v3[i] = i
	// }
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(v3[i])
	// }

	var v4 float64 = 0.555556666666666
	fmt.Println(v4)

	// var v5
	v5 := 777
	fmt.Println(v5)

	var i = 1
	var j = 2
	i, j = j, i

	fmt.Println(i, j)

	var v6 bool //默认False
	fmt.Println(v6)

	v7 := (1 == 1.0)
	fmt.Println(v7)

	v8 := 88888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888
	fmt.Println(v8)
}
