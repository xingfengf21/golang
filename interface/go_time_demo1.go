package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Duration(30))
	fmt.Println(time.Duration(rand.Intn(3) + 300))
	fmt.Println(time.Duration(rand.Intn(30)))
}
