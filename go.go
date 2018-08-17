//变量定义

var x, y, z int
var s, n = "abc", 123
var (
a int
b float32
)
func main() {
n, s := 0x1234, "Hello, World!"
println(x, s, n)
}


data, i := [3]int{0, 1, 2}, 0



func test() (int, string) {
return 1, "abc"
}

_, s := test()
    println(s)
枚举
const (
    Sunday = iota // 0
    Monday // 1，通常省略后续⾏行表达式。
    Tuesday // 2
    Wednesday // 3
    Thursday // 4
    Friday // 5
    Saturday // 6
)


make 和 new

内置函数 new 计算类型⼤⼩，为其分配零值内存，返回指针。
而 make 会被编译器翻译成具体的创建函数，由其分配内存和初始化成员结构，返回对象⽽⾮指针
