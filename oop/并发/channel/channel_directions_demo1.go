package main

import (
	"fmt"
)

//channel方向
//当使用通道作为功能参数时，您可以指定通道是否仅用于发送或接收值

func ping(pings chan<- string, msg string) {
    pings <- msg
}


func pong(pings <- chan string,pongs chan <- string) {
    msg := <-pings 
    pongs <- msg
}

func main() {

    pings := make(chan string ,1)
    pongs := make(chan string ,1)
    ping(pings,"passed message")
    pong(pings, pongs)
    value := <- pongs
    fmt.Println(value)
    fmt.Println()
}
