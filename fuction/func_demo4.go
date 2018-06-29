package main

import "fmt"

func foo1()  {

}

func foo6() interface{}  {
	return foo1
}

func main() {
	v := foo6()
	fmt.Println(v)
}
