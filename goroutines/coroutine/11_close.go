package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go sendData(ch1)

	//main 中读取ch1的通道数据
	for {
		time.Sleep(time.Second)
		v, ok := <-ch1
		// if !ok {
		// 	fmt.Println("已经读完所有数据...", v, ok)
		// 	break
		// }
		// fmt.Println("读取的数据：", v, ok)
		if ok {
			fmt.Println("读取的数据： ", v, ok)
		} else {
			fmt.Println("已经读完所有数据...", v, ok)
			break
		}
	}

	fmt.Println("main over....")
}

func sendData(ch1 chan int) {
	//发送10条数据
	for i := 0; i < 10; i++ {
		ch1 <- i + 1
	}

	close(ch1)
}
