package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// fmt.Println("goroot-->", runtime.GOROOT())
	// fmt.Println("os/platform-->", runtime.GOOS)

	// fmt.Println(runtime.NumCPU())

	// n := runtime.GOMAXPROCS(4)
	// fmt.Println(n)

	//gosched
	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Println("goroutine...", i)
	// 	}
	// }()

	// for i := 0; i < 5; i++ {
	// 	runtime.Gosched() //让出时间片，先让别的goroutine执行
	// 	fmt.Println("main ...", i)
	// }

	go func() {
		fmt.Println("goroutine start...")
		fun()
		fmt.Println("goroutine end...")
	}()

	time.Sleep(3 * time.Second)
}

func fun() {
	defer fmt.Println("defer....")
	// return
	runtime.Goexit()
	fmt.Println("fun函数...")
}
