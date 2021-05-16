package main

import "fmt"

func main() {
	var a chan int
	fmt.Printf("%T %v %p\n", a, a, &a)

	if a == nil {
		fmt.Println("channel 是nil,要make")
		a = make(chan int)
	}

	fmt.Printf("%T %v %p\n", a, a, &a)
	test1(a)

}

func test1(ch chan int) {
	fmt.Printf("%T %v %p\n", ch, ch, &ch)

}
