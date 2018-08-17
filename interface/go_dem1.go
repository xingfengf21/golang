package main

import (
	"fmt"
	"time"
)

func foo1() {
	for {

		fmt.Println("foo1.....")

	}
}

func foo2() {
	for {

		fmt.Println("foo2....")

	}
}

func main() {

	go foo2()
	go foo1()
	time.Sleep(time.Second * 5)
}
