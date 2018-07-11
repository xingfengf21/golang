package main

import (
	"fmt"
)

func printAll(vals []interface{}) { //1
	// var data []string = vals
	// data := vals
	// var interfaceSlice []interface{} = make([]interface{}, len(data))
	for _, val := range vals {
		fmt.Println(val)
	}
}
func main() {
	names := []string{"stanley", "david", "oscar"}
	printAll(names)
}
