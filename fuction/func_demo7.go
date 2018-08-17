package main

import (
	"fmt"
)

func add(x, y int) (z int) {
	defer func() {
		z += 100
	}()
	z = x + y
	return // return z
}

func main() {
	fmt.Println(add(1, 2))
	fn := func() { println("hello world") }
	fn()
}
