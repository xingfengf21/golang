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

/*
make用于内建类型(map,slice和channel)的内存分配,new用于各种类型的内存分配

new本质上和其他语言一样,new(T) 分瞥了零值填充的T类型的内存空间,并返回其地址,即一个*T类型的值,
用Go的术语,它返回了了一个指针,指向新分配的类型T的零值

new返回指针.

func make(t Type, size ...IntegerType) Type

make只能创建slice,map,channel并返回一个有初始值(非零)的T类型,而不是*T
对于slice,map,channel make初始化了内部的数据结构,填充适当的值
make返回初始化后的(非零)值

 */


func main() {
	foo()
	foo3()
	foo2()

}

