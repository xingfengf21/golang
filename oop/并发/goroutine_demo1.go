package main 

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	//分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)
	// wg用来等待程序完成
	//技术加2表示等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	//声明一个匿名函数,并创建一个goroutine
	go func(){
		defer wg.Done()
		for count:=0;count < 3;count++{
			for char:='a'; char <'a'+26;char++{
				fmt.Printf("%c ",char)
			}
		}
	}()

	go func(){
		defer wg.Done() //在函数退出调用时,调用Done通知main函数工作已经完成
		for count:=0;count < 3;count++{
			for char:='A'; char <'A'+26;char++{
				fmt.Printf("%c ",char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\n Terminating Program")
}