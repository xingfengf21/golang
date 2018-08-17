package main

import (
	"os"
	"fmt"
	"net"
)

func main() {

	if len(os.Args) != 2{
		fmt.Println("Usage:",os.Args[0],"host")
		os.Exit(-1)
	}

		service := os.Args[1]
		conn,err := net.Dial("ip4:icmp",service)
		fmt.Println(conn,err)

}

