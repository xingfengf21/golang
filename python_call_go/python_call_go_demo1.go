package main

import "fmt"
import "C"

//export test
//指定那些函数能被外部调用

func test() {
	//test函数循环1000000 次
	for a := 0; a < 1000000; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

}

func main() {
	fmt.Println("hello world")
	test()
}
