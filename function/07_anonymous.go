package main

import "fmt"

//函数名是函数的入口地址
//把函数赋值给一个变量，这不就等于是一个指向函数的指针了吗
//函数的类型包括形参和返回值

//匿名函数就是定义一个没有函数名的函数，用一个函数指针变量指向这个匿名函数，后面用这个变量调用匿名函数

func main() {
	// xxx := hehe

	// fmt.Printf("%T %p %x\n", hehe, hehe, hehe)
	// fmt.Printf("%T %p %d\n", xxx, xxx, &xxx)

	res1 := func(a, b int) int {
		// fmt.Println(a, b)
		return a + b
	}(1, 2)

	fmt.Println(res1)
}

func hehe() {
	fmt.Println("hehe")
}
