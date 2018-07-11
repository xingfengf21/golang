package main

import "fmt"

/*
**用于输出数组元素
 */

//  cannot range over a (type interface {})

// func echoArray(a interface{}) {

// 	for _, v := range a {
// 		fmt.Println(v)
// 	}

// }

func echoArray(a interface{}) {

	for _, v := range a {
		fmt.Println(v)
	}

}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	echoArray(a)
}
