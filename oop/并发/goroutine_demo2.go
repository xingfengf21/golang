package main 

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup

func main(){
	//分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(runtime.NumCPU())
	// wg用来等待程序完成
	//技术加2表示等待两个goroutine
	
	wg.Add(2)
	fmt.Println("creating Goroutines")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\n Terminating Program")
}

func printPrime(prefix string){
	defer wg.Done()
	next:
	for outer:=2;outer<5000;outer++{
		for inner:=2; inner<outer;inner++{
			if outer % inner == 0{
				continue next
			}
		}
		fmt.Printf("%s:%d\n",prefix,outer)
	}
	fmt.Println("Completed",prefix)
}