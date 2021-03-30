package main

import "fmt"

func main() {
	var i1, i2, i3 int32
	fmt.Printf("%v ,%v ,%v\n", i1, i2, i3)

	i4, i5, i6 := 23, 465, 567
	fmt.Printf("%v ,%v ,%v\n", i4, i5, i6)

	//_是一个特殊变量，任何赋予它的值都会被丢弃
	_, b := 11, 22
	fmt.Printf("%v \n", b)

	//go变量声明了必须被使用

}
