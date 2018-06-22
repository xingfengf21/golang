package main

import (
	"fmt"
)

//notifier 是一个定义了通知类型行为的接口
type notifier interface {
	notify()
}

//用户类型
type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func main() {
	//创建一个user类型的值,并发送通知
	u := user{"Bill", "bill@email.com"}
	/*
	  如果使用指针接受着来实现一个接口,那么只有指向那个类型的指针才能实现对应的接口
	  如果使用值接受着实来实现一个接口,南无那个类型的值和指针都能够实现对应的接口
	*/
	sendNotification(&u)
	sendNotification(u)
}

// sendNotification接受一个实现了notifier接口的值
//任何一个实现了notifier接口的值都可以传入sendNotification函数
func sendNotification(n notifier) {
	n.notify()
}
