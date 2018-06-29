package main

import (
	"fmt"
)

/*
1.Go语言最精妙的应该算interface

2.什么是interface, 简单来说,interface是一组method的组合,通过interface来定义对象的一组行为

*/

type Human struct {
	name string
	age int 
	phone string 
}

type Student struct {
	Human
	school string 
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

//Human对象实现Sayhi方法
func (h *Human) Sayhi()  {
	fmt.Printf("Hi,I am %s you can call me on %s\n",h.name,h.phone)
}

//Human对象实现Sing方法

func (h *Human) Sing(lyrics string)  {
	fmt.Println("La la la la...",lyrics)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string)  {
	fmt.Println("Guzzle Guzzle ,,,",beerStein)
}

//Empolyee重载Human的Sayhi方法
func (e *Employee) Sayhi()  {
	fmt.Printf("Hi,I am %s,I work at %s,call me on %s\n",e.name,e.company,e.phone)
}

func (h Human) Fuck() {
	fmt.Println("Fuck you man!")
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney (amount float32){
	s.loan += amount
}

//Employee实现SpendSalary方法

func (e *Employee) SpendSalart(amount float32)  {
	e.money  -= amount
}

//定义interface

type Men interface {
	Sayhi()
	Sing(lyrics string)
	Guzzle(beerStein string)
	Fuck()
}

type YoungChap interface {
	Sayhi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	Sayhi()
	Sing(song string)
	SpendSalary(amout float32)
}

/*
interface值  interface里边能存储什么值
interface 可以存储实现了该接口的任意类型对象
定义一个Men interface类型变量m, m里边可以存Human,Student,Employee的值
任意的类型都实现了空接口,interface{}.包含0个method的infterface,空接口可以存储任何类型的值

interface就是一组抽象方法的集合,它必须由其他非interface类型实现,不能自我实现
go通过interface实现了duck-typing:"当看到一只鸟,走起来像鸭子,游泳像鸭子,叫起来像鸭子,那么这只鸟就可以被称为鸭子"

*/

func foo()  {
	fmt.Println("foo")
}

func main() {
    //var i interface{}
    //i = 3
    //fmt.Println(i)
    //i = foo
    //fmt.Println(i)

    mike := Student{Human{"Mike",23,"222-222-xxx"},"MIT",0.00}
    paul := Student{Human{"Paul",26,"111-222"},"Harvard",100}
    sam := Employee{Human{"Sam",36,"666-888"},"Golang Inc",10000}
    tom := Employee{Human{"Tom",66,"999-888"},"Think",66666666}

    var i Men

    // i能存储Student
    i = &mike
    i.Sayhi()

    i = &paul
    i.Sayhi()
    i = &sam
    i.Sayhi()
    i = &tom
    i.Sayhi()

    var m Men
    m = &mike
    m.Sayhi()
    m.Fuck()
    tom.Fuck()

    //定义了slice Men
    x := make([]Men,3)

    // T这三个都是不同类型的元素,但是他们都实现了interface同一个接口
    x[0],x[1],x[2] = &paul,&sam,&mike
	for _,value := range x{
		value.Sayhi()
	}
}


















