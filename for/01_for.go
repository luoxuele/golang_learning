package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//相当于while
	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println(i)

	//while(true)
	for {
		fmt.Println("true")
	}
}
