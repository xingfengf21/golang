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

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func main() {
	//创建一个user类型的值,并发送通知
	u := user{"Bill", "bill@email.com"}
	sendNotification(&u)
}

// sendNotification接受一个实现了notifier接口的值

func sendNotification(n notifier) {
	n.notify()
}
