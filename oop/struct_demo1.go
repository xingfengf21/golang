package main

import "fmt"

type Rect struct{
	x,y float64
	width,height float64 
}

func (r *Rect) Area() float64{
	return r.width * r.height
}

//全局创建函数 ---  构造函数
func NewRect(x,y,width,height float64) *Rect{
	return &Rect{x,y,width,height}
}

func main(){
	//结构体的初始化
	rect1  := new(Rect)
	rect2 := &Rect{1,2,10,20}
	rect3 := &Rect{width:100,height:200}
	fmt.Println(rect1.Area())
	fmt.Println(rect2.Area())
	fmt.Println(rect3.Area())
	fmt.Println()

	newRect := NewRect(1,2,3,4)
	fmt.Println(newRect.Area())
}