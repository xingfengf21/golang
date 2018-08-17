package main

import (
	"fmt"
)

func main() {
	type data struct {
		a int
	}
	var d = data{666}
	var p *data
	p = &d
	fmt.Printf("%p, %v\n", p, p.a)

	var i int = 0
	var pi *int = &i
	fmt.Println(i)
	*pi++
	fmt.Println(i)
}
