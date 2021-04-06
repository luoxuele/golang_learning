package main

import "fmt"

func main() {

	//1. 创建channel
	ch := make(chan int, 3)

	ch <- 10
	ch <- 21
	ch <- 31

	a := <-ch
	fmt.Println(a)

	<-ch
	// <-ch
	c := <-ch
	fmt.Println(c)
	fmt.Println(ch, cap(ch), len(ch))

	ch2 := make(chan int, 1)
	ch2 <- 34
	ch2 <- 33 //all goroutines are asleep - deadlock!

}
