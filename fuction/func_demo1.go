package main

import (
	"fmt"
	"reflect"
)

//变参
//arg是int类型的切片,[]int
func foo(arg ...int)  {
	fmt.Println("Type",reflect.TypeOf(arg))
	for _,n :=range arg{
		fmt.Println(n)
	}
}

func main() {
	//lists:= []int{1,2,3,4}
	foo()
	foo(1,2,3,4)
}
