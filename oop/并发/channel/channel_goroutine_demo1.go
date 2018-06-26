package main

import "fmt"

var counts int = 0

func Count(i int, ch chan int) {
	fmt.Println(i, "WriteStart")
	ch <- 1
	fmt.Println(i, "WriteEnd")
	fmt.Println(i, "end", "and echo", i)
	counts++
}

func main() {
	chs := make([]chan int, 3)
	for i := 0; i < 3; i++ {
		chs[i] = make(chan int)
		fmt.Println(i, "ForStart")
		go Count(i, chs[i])
		fmt.Println(i, "ForEnd")
	}

	fmt.Println("Start debug")
	for num, ch := range chs {
		fmt.Println(num, "ReadStart")
		<-ch
		fmt.Println(num, "ReadEnd")
	}

	fmt.Println("End")

	//为了使每一步数值全部打印
	for {
		if counts ==3 {
			break
		}
	}
}
