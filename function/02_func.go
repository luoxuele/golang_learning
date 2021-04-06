package main

import "fmt"

func main() {

	fmt.Println(getSum(100))
}

func getSum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}
