package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		fmt.Println("sub goroutine start....")
		data := <-ch1
		fmt.Println("sub goroutine end ... data:", data)
	}()

	time.Sleep(5 * time.Second)
	ch1 <- 10
	fmt.Println("main over....")

	// ch := make(chan int)
	// ch <- 100 //阻塞
}
