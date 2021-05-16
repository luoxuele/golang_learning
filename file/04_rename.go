package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Rename("./xxx.txt", "./hhh.txt")

	if err != nil {
		fmt.Println(err)
	}
}
