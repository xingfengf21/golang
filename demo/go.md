
why go?

1.并行与分布式支持

2.软件工程支持

3.编程哲学的重塑 面向对象  --- 什么才是最佳的编程实战


多核化和集群化是互联网典型特征

并发执行的"执行体": 进程 ,线程,协程(coroutine)


go的特点:
1.语言级别支持协程

编程哲学:

a. Go反对函数和运算符重载
java/c++ 允许同名函数 只要参数列表不同就行

重载带来了哪些负担?

b.GO支持类,类成员方法,类的组合(组合方式提供继承)

反对:继承,虚函数,虚函数重载
反对的原因是什么?


c.GO放弃了构造函数和析构函数

"如果一个特性并不对解决问题有显著的价值,Go就不提供它"


d.GO的接口

与其他语言的区别是它的非入侵性

c++,java,c#中,为了实现一个接口,需要从该接口中继承

java

class Foo implement IFoo{
    
}

c++

class Foo : public IFoo{
    
}

IFoo * foo = new Foo;


go,实现类的时候无需从接口中派生

type Foo struct{
    
}

# 只要Foo实现了接口IFoo要求的所有方法,就实现了该接口,可以进行赋值

var foo IFoo = new(Foo)


e.函数式编程   匿名函数和闭包


f.Erlang语言的面向消息编程思想,支持goroutine和通道,推荐使用消息而不是共享内容进行并发编程


语言特性:

自动垃圾回收

更丰富的内置类型 

字典类型 map
数组切片 --- c++ vector非常类似

函数多返回值  
错误处理
匿名函数和闭包 

GO语言中,所有的函数也是值类型, 什么是值类型?

定义了一个名为f的匿名函数

f := func(x, y int) int {
    return x + y
}

类型和接口

c++中的类型和接口

interface IFly
{
    virtual void Fly() = 0;
};

// 实现类

class Bird : puclic IFly
{
    puclic:
    Bird(){};
    virtual ~Bird();

public:
    void Fly(){
        //
    }
}

main{
    IFly * pFly = new Bird();
    pFly->Fly();
    delete pFly;
}

在实现一个接口之前必须先定义该接口,并且将类型和接口紧密绑定
即接口的修改会影响到所有实现了该接口的类型


Go的接口

type Bird struct{
    ...
}

func (b *Bird) Fly(){
    //以鸟的方式飞行
}



//在实现Bird类型时,完全没有任何IFly的信息

type IFly interface{
    Fly()
}

func main(){
    var fly IFly = new(Bird)
    fly.Fly()
}

并发编程
反射
语言交互性



1.
package main

要生成Go可执行程序,必须建立main的包,并包含main函数


main函数不能带参数,也不能定义返回值 命令行参数 os.Args

不能包含源码中没有用到的包



Go函数定义

func 函数名(参数列表)(返回值列表){
    //函数体
}

func Compute(value1 int ,value2 float64) (result float64,err error){
    //函数体
}

函数返回时,没有被明确赋值的返回值会被设置为默认值

resutl默认0.0 err 默认为nil


注释 : 同 c++

语句结束没有分号






Go字符串操作

+
len()
s[i]

字符串遍历

