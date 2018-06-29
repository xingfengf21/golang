package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := errors.New("IOError")
	if  err != nil {
		fmt.Println(err)
	}
}
