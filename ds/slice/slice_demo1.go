package main

import "fmt"

func main()  {

	var array = [10]byte{'a','b','c','d','e','f'}
	fmt.Println(array)
	fmt.Println(array[:3])
	fmt.Println(array[3:10])
}
