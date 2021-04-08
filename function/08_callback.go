package main

import "fmt"

//回调，就是把一个函数当作参数传给另一个函数，然后在另一个函数中调用

func main() {
	// fmt.Println(add(3, 4))
	// fmt.Println(sub(3, 4))
	a := oper(3, 4, add)
	fmt.Println(a)
	a = oper(44, 33, sub)
	fmt.Println(a)
}

func oper(a, b int, fun func(int, int) int) int {
	return fun(a, b)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
