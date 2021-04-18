package main

import "fmt"

func main() {

	fmt.Println(sum(100))
}

func sum(len int) int {
	sum := 0
	for i := 1; i <= len; i++ {
		sum += i
	}
	return sum
}
