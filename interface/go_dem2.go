package main

import (
	"fmt"
)

/*类型通过实现方法来实现接口，却不必要显示的声明。所以没有关键字 implements 。这是隐式接口*/

type Abser interface {
	Abs() float64
}
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return v.X*v.X + v.Y*v.Y
}
func (v Vertex) Scale() float64 {
	return v.X + v.Y
}

func (v Vertex) Mult() float64 {
	return v.X * v.Y
}
func main() {
	var a Abser
	f := Vertex{3, 4}
	a = &f

	fmt.Println(f.Abs())
	fmt.Println(f.Scale())
	fmt.Println(a.Abs())
	fmt.Println(a.Mult())
	//fmt.Println(a.Scale())
}
