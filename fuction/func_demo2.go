package main

import "fmt"

func add1(a int)int  {
	a = a + 1
	return  a
}

func add2(a *int)int  {
	*a += 1
	return *a
}

//传值和传指针
/*
传指针的好处
1.多个函数可以操作同一个对象
2.传指针只是传内存地址,比较轻量级,可以用指针传递体积大的结构体
3.string ,slice,map实现机制类似指针.所以可以直接传递
*/


func main() {
	x := 3
	fmt.Println("x = ",x)
	x1 := add1(x)
	fmt.Println("x+1 = ",x1)
	fmt.Println("x = ",x)

	y := 4
	fmt.Println("y = ",y)
	fmt.Println("y+1 = ",add2(&y))
	fmt.Println("y =",y)

}