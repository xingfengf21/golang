package main 

import "fmt"

type user struct{
	name string 
	email string 

}

//方法能给用户定义的类型添加新的行为,方法实际上是函数
// 只是在声明时,在关键字func和方法名之间增加了一个参数

//值接收者, 方法notify会接收到u的一个副本
func (u user) notify(){
	fmt.Printf("Sending User Email To %s<%s>\n",u.name,u.email)
}

//指针接收者
func (u *user) changeEmail(email string){
	u.email = email
}

func main(){
	
	bill:=user{"Bill","bill@email.com"}
	bill.notify()

	lisa := &user{"Lisa","lisa@email.com"} //声明一个名为lisa的指针变量,并使用指定的名称和邮箱初始化
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	(&bill).changeEmail("bill@&bill_newdomain.com")
	bill.notify()
	fmt.Println("bill.email",bill.email)
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
	(*lisa).notify() //指针被解引用为值,notify操作的是一个从lisa指针指向的值的副本

}