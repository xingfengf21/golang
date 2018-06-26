package main

import (
	"fmt"
)

/*
1. channel是Go在语言层次提供的goroutine的通信方式
2.可以使用channel在多个goroutine之间传递消息
3.channel是进程内的通信方式
4.channel是类型相关的,一个channel只能传递一种类型的值
5. var chanName chan ElementType  声明一个传递类型为ElementType的channel

var m map[string] chan bool声明一个map,元素类型是bool的channel

6.channel的定义:
 ch := make(chan int)
 声明并初始化了一个int型的名为ch的channel

*/

func Count(ch chan int, i int) {
	ch <- i //数据写入到channel
	fmt.Println("counting")

}

func main() {

	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i], i)
	}
	for _, ch := range chs {
		value := <-ch //从channel中读取数据
		fmt.Println(value)
	}
}
