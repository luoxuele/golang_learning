package main

import "sync"

var wg sync.WaitGroup

func main() {
	wg.Add(4)
	go xxx()
	go xxx()
	go xxx()
	go xxx()

	wg.Wait()
}

func xxx() {
	defer wg.Done()
	for {
		//fmt.Println("1")
	}
}
