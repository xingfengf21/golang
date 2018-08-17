package main

import (
	"fmt"
	"reflect"
)

func main() {
	var circle float64 = 6.28
	value := reflect.ValueOf(circle)
	fmt.Println(value.CanSet()) // false
	fmt.Println(value)          // 6.28

	value2 := reflect.ValueOf(&circle)
	fmt.Println(value2.CanSet()) // false
	fmt.Println(value2)          // 0xc420086008

	value3 := reflect.ValueOf(&circle).Elem()
	fmt.Println(value3.CanSet()) // true
	fmt.Println(value3)          // 6.28

	value3.SetFloat(3.14)
	fmt.Println(circle) // 3.14

}
