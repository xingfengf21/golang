package main

import "fmt"

func main() {
	var arr [10]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)

	a := [...]int{1,2,3}
	fmt.Println(a)
}
