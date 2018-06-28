package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string,3)

func main() {
	syncChan1 := make(chan struct{},1)
	syncChan2 := make(chan  struct{},2)
	go receive(strChan,syncChan1,syncChan2)
	go send(strChan,syncChan1,syncChan2)
	<-syncChan2
	<-syncChan2
}

func receive(strChan <- chan string,
	syncChan1 <- chan struct{},
	syncChan2 chan <-   struct{})  {
	<-syncChan1
	fmt.Println("Received a sync signal and wait a second...[receiver]")
	time.Sleep(time.Second)
	for elem := range strChan{
		fmt.Println("Received:",elem,"[receiver]")
	}
	fmt.Println("Stopped.[receiver]")
	syncChan2 <-struct{}{}
}

//syncChan2 chan<-  struct{}  往管道里写
//syncChan1 <- chan struct{} 从管道里读

func send(strChan chan <-  string,
	syncChan1 chan <-  struct{},
	syncChan2 chan<-  struct{})  {
	for _,elem := range []string{"a","b","c","d"}{
		strChan <- elem
		fmt.Println("Send",elem,"[sender]")
		if elem == "c"{
			fmt.Println("Send a sync signal,[sender]")
			syncChan1 <- struct{}{}
		}
	}

	fmt.Println("Wait 2 seconds ...[sender]")
	time.Sleep(time.Second * 2)
	close(strChan)
	syncChan2<- struct{}{}
}
