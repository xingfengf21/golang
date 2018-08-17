/*
字符串是不可变值类型，内部⽤用指针指向 UTF-8 字节数组。
• 默认值是空字符串 ""。
• ⽤用索引号访问某字节，如 s[i]。
• 不能⽤用序号获取字节元素指针，&s[i] ⾮非法。
• 不可变类型，⽆无法修改字节数组。
• 字节数组尾部不包含 NULL。
rune int32别名


修改字符串:
可先将其转换成 []rune 或 []byte，完成后再转换为 string。⽆无论哪种转
换，都会重新分配内存，并复制字节数组

遍历字符串:


*/

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%T\n", 'a')
	var c1, c2 rune = '\u6211', '们'
	fmt.Println("666", c1 == '我', string(c2) == "\xe4\xbb\xac")

	s := "abcd"
	bs := []byte(s)

	s2 := string(bs)
	fmt.Println("s2", s2)

	bs[1] = 'B'
	//bs[2] = "C" // cannot use "C" (type string) as type byte in assignment
	//bs[5] = '5' //panic: runtime error: index out of range
	fmt.Println(string(bs))
	u := "电脑"
	us := []rune(u)
	us[1] = '话'
	fmt.Println(string(us))

	s1 := "abc汉字"
	for i := 0; i < len(s1); i++ { // byte
		fmt.Printf("%c,", s1[i])
	}
	fmt.Println()
	for _, r := range s1 { // rune
		fmt.Printf("%c:", r)
	}
}
