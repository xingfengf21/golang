package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

var mapChan = make(chan map[string]*Counter,1)

func main() {
	syncChan := make(chan struct{},2)
	go func() {

		for{
			if elem,ok := <- mapChan;ok{
				fmt.Println("receive",*elem["count"],"[receiver]")
				counter := elem["count"]
				counter.count++
			}else{
				break
			}
		}
		fmt.Println("stopped. [receiver]")
		syncChan <- struct{}{}
	}()

	// 发送
	go func() {
		countMap := map[string]*Counter{
			"count":&Counter{},
		}
		for i:=0; i<5; i++ {
			mapChan <- countMap
			time.Sleep(time.Microsecond)
			fmt.Printf("The count map:%v [sender]\n",*countMap["count"])
		}
		time.Sleep(time.Second * 2)
		close(mapChan)


		syncChan<- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

