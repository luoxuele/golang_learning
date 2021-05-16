package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("./xxx.txt")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Printf("%T\n", fileInfo)
	fmt.Println(fileInfo)

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.ModTime())

}
