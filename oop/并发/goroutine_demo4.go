package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {

	lock := &sync.Mutex{}
	for i := 0; i < 88; i++ {
		go Count(lock)
	}
	// for 循环用于不断检测counter的值,也需要加锁
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 88 {
			break
		}
	}
}
