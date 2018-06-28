package  main

import "fmt"

func main() {

		dadaChan := make(chan int,5)
		syncChan1 := make(chan struct{},1)
	    syncChan2 := make(chan struct{},2)
		go func() {
			<- syncChan1
			for{
				if elem,ok := <-dadaChan;ok {
					fmt.Println("receive", elem, "[receiver]")
				}else{
					break
				}
			}
			fmt.Println("Done receiver")
			syncChan2<- struct{}{}
		}()

	    go func() {
			for i:=0 ;i<5 ;i++  {
				dadaChan<- i
				fmt.Println("Send",i)
			}
			close(dadaChan)
			//close(dadaChan) panic: close of closed channel
			fmt.Println("Done sender")
			syncChan1<- struct{}{}
			syncChan2 <- struct{}{}
		}()
	    <-syncChan2
		<-syncChan2
}