package main

import (
	"fmt"
)

func main() {

	//定义一个数组

	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]
	fmt.Println("Elements of myArray:")
	for _, v := range myArray {
		fmt.Println(v, " ")
	}
	fmt.Println("\nElements of mySlice")

	for _, v := range mySlice {
		fmt.Println(v, " ")
	}
	fmt.Println()

	fmt.Println(myArray[:])
	fmt.Println(myArray[1:5])
	mySlice2 := []int{11, 12, 13}
	myslice3 := append(mySlice, mySlice2...)
	fmt.Println(myslice3)
	fmt.Println(len(myslice3))
	fmt.Println(cap(myslice3))

	oldSlice := []string{"a", "b", "c", "d"}
	newSlice := oldSlice[:3]
	fmt.Println(newSlice)
}
