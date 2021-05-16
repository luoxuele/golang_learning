package main

import (
	"fmt"
	"time"
)

func main() {

	//创建并启动子goroutine,执行printnum函数
	go printNum()

	for i := 0; i < 100; i++ {
		fmt.Println("main goroutince :A\t", i)
	}

	fmt.Println("main ...over")
	time.Sleep(1 * time.Second)
}

func printNum() {
	for i := 0; i < 100; i++ {
		fmt.Printf("sub goroutine :%d \n", i)
	}

	fmt.Println("sub ... over")
}
