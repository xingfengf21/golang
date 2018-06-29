package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human
	speciality string
}

func (self Student) GetName() string  {
	return self.name
}

func (self *Student) SetName(name string) {
	self.name = name
}

func main() {
    //初始化一个学生
    mark := Student{Human{"Mark",25,120},"Computer Science"}
    //访问属性
    fmt.Println(mark.name)
    fmt.Println(mark.weight)
    fmt.Println(mark.speciality)
    mark.SetName("Mark6")
    fmt.Println(mark.GetName())
    mark.name = "MMark"
    fmt.Println(mark.name)
}
