理解 Go interface 的 5 个关键点

1.type I interface {
    Get() int
}

interface 是一种具有一组方法的类型，这些方法定义了 interface 的行为

接口是没有属性的对象?


2.interface 变量存储的是实现者的值

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

3. 空的 interface

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	names := []string{"stanley", "david", "oscar"}

	// 对接口数组赋值前，必须多一个类型转换操作

	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	PrintAll(vals)
}



