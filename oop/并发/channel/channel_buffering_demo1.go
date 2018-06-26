package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 3)
	messages <- "string1"
	messages <- "string2"

	for i := 0; i < 3; i++ {
		value := <-messages
		fmt.Println(value)
	}
}
