package main

import (
	"time"
)

func main() {
	for i := 0; i < 100000; i++ {
		go println(i)
	}

	time.Sleep(time.Microsecond)
}
