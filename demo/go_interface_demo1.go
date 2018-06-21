package main

import (
	"fmt"
)

//定义了一个接口Phone，接口里面有一个方法call()

/* 定义接口
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}
*/
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

/* 定义结构体
type struct_name struct {
   //variables
}
*/
type IPhone struct {
}

/* 实现接口方法 */
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	/*
		定义了一个Phone类型变量，并分别为之赋值为NokiaPhone和IPhone。然后调用call()方法
	*/
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
	
}
