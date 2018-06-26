package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe() //指定stdin连接StdinPipe,然后操作stdin就实现了对command的参数传递
	if err != nil {
		fmt.Println(err)
	}
	_, err = stdin.Write([]byte("tmp.txt")) //字节切片
	if err != nil {
		fmt.Println(err)
	}
	stdin.Close()
	cmd.Stdout = os.Stdout //终端标准输出tmp.txt
	cmd.Start()
}
