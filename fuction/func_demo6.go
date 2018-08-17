package main

import (
	"fmt"
)

func test(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}
func main() {
	println(test("sum: %d", 1, 2, 3))
	s1 := []int{1, 2, 3} //使⽤用 slice 对象做变参时，必须展开
	println(test("sum: %d", s1...))
}
