package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	ra, _ := ioutil.ReadFile("win.ini")
	fmt.Printf("%s", ra)
	s := "heLLo worLd fuck you man "
	ts := strings.Title(s)
	fmt.Printf("%q\n", ts) // "HeLLo WorLd"

}
