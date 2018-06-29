package main

import "fmt"

//map[keyType]valueType

//map的读取和设置和slice一样,通过key操作,slice的key只能是int,map多了很多类型

func main() {
	//声明一个key是字符串,值是int的字典,这种声明需要在使用之前用make初始化
	var numbers map[string]int
	numbers = make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	numbers["four"] = 4
	fmt.Println(numbers)

	delete(numbers,"one1")
	fmt.Println(numbers)

	//map 是引用类型 两个map指向同一个底层,一个改变另一个也改变

	numbers2 := numbers
	numbers["one"] = 666
	fmt.Println(numbers2)
	fmt.Println(numbers)

}
