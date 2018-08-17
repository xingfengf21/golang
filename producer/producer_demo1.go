package main

import (
	"fmt"
	"time"
)

func consumer(cname string, ch chan int) {
	for i := range ch {
		fmt.Println("consumer--", cname, ":", i)
	}
	fmt.Println("ch closed.")
}

func producer(pname string, ch chan int) {
	for i := 0; i < 4; i++ {

		fmt.Println("producer--", pname, ":", i)
		ch <- i
	}
}

func main() {
	//用channel来传递"产品", 不再需要自己去加锁维护一个全局的阻塞队列
	data := make(chan int)
	go producer("生产者1", data)
	go producer("生产者2", data)
	go consumer("消费者1", data)
	go consumer("消费者2", data)

	time.Sleep(1 * time.Second)
	close(data)
	time.Sleep(1 * time.Second)
}
