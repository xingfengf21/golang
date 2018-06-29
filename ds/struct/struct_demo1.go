package main

import "fmt"

type person struct {
	name string
	age int
}


func foo()  {
	var P person
	P.name = "康东"
	P.age = 29
	fmt.Printf("The person's name is %s\n",P.name)

	// 安装定义顺序,初始化值
	P2 := person{"Tom",25}
	fmt.Println(P2)

	P3 := person{age:24,name:"Alex"}
	fmt.Println(P3)
}

func main() {
	foo()
}
