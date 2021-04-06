package main

import "fmt"

//自定义类型
type myInt int
type myFn func(int, int) int

//类型别名
type myFloat = float64

func main() {

	var a myInt = 10
	fmt.Printf("%T %v\n", a, a)

	var b myFloat = 12.3
	fmt.Printf("%T %v\n", b, b)

}
