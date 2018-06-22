package main 

//值语义和引用语义

import std "fmt"
import "reflect"

func main(){
	var a = [3] int{1,2,3}
	var b = a 
	b[1]++
	std.Println(a,b)
	std.Println()
	var c = [3] int{1,2,3}
	var d = &c // 数组内容的引用 d的类型是 *[3] int 类型, c的类型是[3] int类型
	c[1]++ 
	std.Println(c,d,*d)
	std.Println(reflect.ValueOf(d))
	std.Println(reflect.ValueOf(c))

}