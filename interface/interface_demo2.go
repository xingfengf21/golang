package main

import (
	"fmt"
)

//interface变量存储的是实现者的值

//1 定义接口 I
type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

// S实现了I接口的两个方法

func (s S) Get() int {
	return s.Age
}
func (s *S) Set(age int) {
	s.Age = age
}

//定义了一个函数 f 参数类型是 I，S 实现了 I 的两个方法就说 S 是 I 的实现者
//3
func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	s := S{}
	// f(&s) //4  f(&s) 就完了一次 interface 类型的使用

	s.Age = 66
	var i I //声明 i
	i = &s  //赋值 s 到 i
	fmt.Println(i.Get())

	// 类型断言
	if t, ok := i.(*S); ok {
		fmt.Println("s implements I", t)
	}

}
