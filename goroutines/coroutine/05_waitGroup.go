package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //创建同步等待组对象
func main() {
	wg.Add(2)

	go fun1()
	go fun2()

	fmt.Println("main 进入阻塞状态，等待wg中的子goroutine结束...")
	wg.Wait()
	fmt.Println("main ...解除阻塞")

}

func fun1() {
	for i := 1; i < 100; i++ {
		fmt.Println("fun1()...", i)
	}

	wg.Done()
}

func fun2() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Println("fun2()...", i)
	}
}
