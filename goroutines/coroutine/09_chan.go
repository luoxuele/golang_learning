package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var ch1 chan bool
	ch1 = make(chan bool)

	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("goroutine--- ", i)
		}

		ch1 <- true
		wg.Done()
	}()

	go func() {
		data := <-ch1
		fmt.Println("goroutine2 read 成功....", data)
		wg.Done()

	}()

	fmt.Println("main ... start...")
	// data := <-ch1
	// fmt.Println("main end....", data)
	wg.Wait()
	fmt.Println("main end....")

}
