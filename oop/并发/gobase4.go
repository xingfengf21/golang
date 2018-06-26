package main

import (
	"fmt"
	"time"
)

func foo1() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func() {
			fmt.Printf("Hello,%s\n", name)
		}()
	}
	time.Sleep(time.Microsecond)
}

func foo2() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello,%s\n", who)
		}(name)
	}
	time.Sleep(time.Microsecond)
}

func main() {
	foo1()
	fmt.Println("-------------------------")
	foo2()
}
