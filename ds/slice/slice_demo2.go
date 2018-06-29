package main

import (
	"fmt"
	"reflect"
)

//空接口可以指向任何数据对象，所以可以使用interface{}定义任意类型变量，同时interface{}也是类型安全的

func main() {

	f(2, "Go", 8, "language", 'a', false, 'A', 3.14,struct{}{},[]interface{}{})

}

func f(args ...interface{}) {

	fmt.Println("type",reflect.TypeOf(args))  // args 是空接口类型的切片  []interface {}

	var num = make([]int, 0, 6)

	fmt.Println("num type:",reflect.TypeOf(num),num,len(num),cap(num))

	var str = make([]string, 0, 6)

	var ch = make([]int32, 0, 6)

	var other = make([]interface{}, 0, 6)

	for _, arg := range args {

		switch v := arg.(type) {

		case int:

			num = append(num, v)

		case string:

			str = append(str, v)

		case int32:

			ch = append(ch, v)

		default:

			other = append(other, v)

		}

	}

	fmt.Println(num)

	fmt.Println(str)

	fmt.Println(ch)

	fmt.Println(other)

}
