package main 

import (
	"bytes"
	"fmt"
	"io"
	"os"
)
/*
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
Fprintf接受一个io.Writer类型的接口值作为第一个参数
b.Buffer类型的指针实现了io.Writer接口,所以可以将b传入fmt.Fprintf
类型实现了哪写接口怎么看?
*/
func main(){
	var b bytes.Buffer

	b.Write([]byte("Hello"))
	fmt.Fprintf(&b,"World!")
	io.Copy(os.Stdout,&b)
}