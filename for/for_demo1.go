package main

import (
	"fmt"
)

//https://blog.csdn.net/benben_2015/article/details/78904860

/*
go语言中的字符串实际上是类型为byte的只读切片。
或者说一个字符串就是一堆字节。这意味着，当我们将字符存储在字符串中时，实际存储的是这个字符的字节。一个字符串包含了任意个byte，它并不限定Unicode，UTF-8或者任何其他预定义的编码
go用byte类型表示字符串
*/

func main() {
	s := "abc"
	for i, n := 0, len(s); i < n; i++ { // 常⻅见的 for 循环，⽀支持初始化语句。
		println(s[i])
	}
	// n := len(s)
	// for n > 0 { // 替代 while (n > 0) {}
	// 	println(s[n]) // 替代 for (; n > 0;) {}

	// 	n--
	// }
	// for { // 替代 while (true) {}
	// 	println(s) // 替代 for (;;) {}
	// }
	for i, v := range s {
		println(i, v)
		//fmt.Printf("%q\n", v)
		fmt.Println(string(v))
	}

	var i int
	for i < len(s) {
		fmt.Printf("s[%d] = %c \n", i, s[i])
		i++
	}
}
