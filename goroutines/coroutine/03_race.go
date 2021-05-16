package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1

	go func() {
		fmt.Println("goroutine...", a)
	}()

	a = 3
	time.Sleep(1 * time.Second)
	fmt.Println("main ...", a)
}

//go run -race 03_race.go
