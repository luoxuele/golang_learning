package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("01_Open.go")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	// bufio 读文件
	reader := bufio.NewReader(file)
	for {
		//一行一行的读
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("---")
			fmt.Println(err)
			return
		}
		fmt.Println(str)
	}

}
