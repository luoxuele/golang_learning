package main

import "fmt"

func main() {
	//array slice map

	//array
	// var arr[]type
	//数组是值类型，作为参数传入函数的时候，传入的是该数组的副本

	var arr [10]int
	arr[0] = 111
	arr[9] = 222
	fmt.Println(arr)

	a := [3]int{1, 2, 3}
	b := [10]int{1, 2, 3}
	c := [...]int{4, 5, 6} //自动计算数组长度
	fmt.Println(a, b, c)

	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(doubleArray, easyArray)

	//slice
	//一个引用类型，总是指向一个底层array
	// var fslice []int
	fslice := []byte{'1', 'a', 'c', 'd'}
	fmt.Printf("%T %v\n", fslice, fslice)

	//slice 内置函数
	//len cap append copy
	fslice = append(fslice, 'f')
	len := len(fslice)
	cap := cap(fslice)

	fmt.Println(len, cap)

	// var array [10]int
	// slice := array[2:4]
	// fmt.Println(slice, len(slice), cap(slice))

	//map 引用类型
	var numbers map[string]int
	numbers = make(map[string]int)
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3
	delete(numbers, "one")
	fmt.Println(numbers)

	v, index := numbers["one"]
	fmt.Println(v, index)

	//map特点
	//不是thread-safe

	//make new 操作
	// 	make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。

	// 内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：

	// new返回指针。

	// 内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。

	// make返回初始化后的（非零）值。

}
