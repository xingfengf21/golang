package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var a int = 30

	// v := reflect.ValueOf(a) //返回Value类型对象，值为30
	// t := reflect.TypeOf(a)  //返回Type类型对象，值为int
	// fmt.Println(v)
	// fmt.Println(t)

	// v = reflect.ValueOf(&a) //返回Value类型对象，值为&a，变量a的地址
	// t = reflect.TypeOf(&a)  //返回Type类型对象，值为*int
	// fmt.Println(v)
	// fmt.Println(t)

	//通过使用reflect.ValueOf和reflect.TypeOf将接口类型变量分别转换为反射类型对象value和type

	var a int = 30
	value := reflect.ValueOf(&a)  //返回Value类型对象，值为&a，变量a的地址
	t := value.Interface().(*int) //类型断言，断定v1中type=*int
	fmt.Printf("%T %v\n", t, t)   // *int 0xc420086008
	fmt.Println(*t)               // 30

	cir := 6.28
	value2 := reflect.ValueOf(cir)
	t2 := value2.Interface().(float64)
	fmt.Printf("%T %v\n", t2, t2) // float64 6.28
	fmt.Println(t2)               // 6.28

	t3 := value2.Interface().(int) //error : interface conversion: interface {} is float64, not int
	fmt.Println("t3", t3)

}
