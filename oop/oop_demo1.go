package main

import "fmt"

type Integer int 

//为类型添加方法

func (a Integer) Less(b Integer) bool{
	return a < b
}

//修改对象时,使用指针
func (a *Integer) Add(b Integer){
    *a += b
}

//go的类型基于值传递和C一样,要想修改变量的值,只能传递指针
func (a Integer) Add2(b Integer){
	a += b
}

func main(){
	var a Integer = 1 
	if a.Less(2){
		fmt.Println(a,"Less 2")
	}
	a.Add(2)
	fmt.Println(a)
	a.Add2(1000)
	fmt.Println(a)
}
