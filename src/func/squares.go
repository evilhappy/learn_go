package main

import "fmt"

/*
匿名函数

拥有函数名的函数只能在包级语法块中被声明， 通过函数字面量（function literal），我们可绕过这一限制， 在任何表达式中表示一个函数值。
函数字面量的语法和函数声明相似，区别在于func关键字后没有函数名。
函数值字面量是一种表达式，它的值被成为匿名函数（ anonymous function）

函数squares返回另一个类型为 func() int 的函数。
对squares的一次调用会生成一个局部变量 x 并返回一个匿名函数。
每次调用时匿名函数时， 该函数都会先使x的值加1， 再返回x的平方。
第二次调用squares时， 会生成第二个x变量， 并返回一个新的匿名函数。
新匿名函数操作的是第二个x变量

squares的例子证明，函数值不仅仅是一串代码，还记录了状态。
在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引用。
这就是函数值属于引用类型和函数值不可比较的原因。
Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包。

*/

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main()  {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
