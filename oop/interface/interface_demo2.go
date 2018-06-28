package main

import "fmt"

type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func(s S) Get()int {
	return s.Age
}

func (s S) GetAge() int {
	fmt.Println("Age:",s.Age)
	return s.Age
}

func(s *S) Set(age int) {
	s.Age = age
}

//3
func f(i I){
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	s := S{}
	s.Age = 12
    s.GetAge()
	f(&s)  //4

	var s1 S
    s1.Age = 66
    s1.GetAge()

    //s1 := S{"Age":111}

}
