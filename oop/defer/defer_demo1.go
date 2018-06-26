package main

import (
	"fmt"
)

func foo() {
	defer fmt.Println("end")
	fmt.Println("foo start...")

}

func main() {

	foo()
}
