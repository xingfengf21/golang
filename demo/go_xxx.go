package main 

import (
    "fmt"
    "time"
)

func foo() {
    for {
        fmt.Println(6)
    }
}

func main() {
    go foo()

    for i := 0; i < 1000; i++ {
        time.Sleep(100)
    }
}