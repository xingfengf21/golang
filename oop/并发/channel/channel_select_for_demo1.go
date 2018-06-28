package main

import "fmt"

func Send(intChan chan <- int,data int)  {
	intChan <- data
}

func main() {
	intChan := make(chan int,10)
	for i:=0 ;i < 10 ;i++  {
		Send(intChan,i)
	}
	close(intChan)
	syncChan := make(chan struct{},1)
	go func() {
		for{
			select {
			case e,ok := <-intChan:
				if !ok{
					fmt.Println("End")
					break
				}
				fmt.Printf("Received:%v\n",e)
			}

	}
		syncChan <- struct{}{}
	}()
	<- syncChan
}


