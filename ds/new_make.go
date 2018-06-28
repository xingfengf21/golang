package main
import (
	"fmt"
	"sync"
)

//对于引用类型的变量，我们不光要声明它，还要为它分配内容空间
func foo()  {
	var i *int
	i=new(int)
	*i=10
	fmt.Println(*i)

}

func foo2()  {
	var i *int
	*i=10
	fmt.Println(*i)
}

func foo3()  {
	u:=new(user)
	u.lock.Lock()
	u.name = "张三"
	u.lock.Unlock()
	fmt.Println(u)
}

type user struct {
	lock sync.Mutex
	name string
	age int
}

func main() {
	foo()
	foo3()
	foo2()

}

