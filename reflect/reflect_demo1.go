package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Id   int
    Name string
    Age  int
}

func (u User) Hello(name string) {
    fmt.Println("Hello", name, "My name is", u.Name)
}

func main() {
    u := User{1, "OK", 12}
    v := reflect.ValueOf(u)
    fmt.Println("vvvv", v.MethodByName("Hello"))
    mv := v.MethodByName("Hello")
    args := []reflect.Value{reflect.ValueOf("JOE")}
    fmt.Println("args",args)
    mv.Call(args)
}