package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go sendData(ch1)

	for v := range ch1 {
		time.Sleep(time.Second)
		fmt.Println("data: ", v)
	}

}

func sendData(ch1 chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}
