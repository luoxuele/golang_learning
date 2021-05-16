package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./aa.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer file.Close()

	bs := make([]byte, 4, 4)
	// n, err := file.Read(bs)
	// fmt.Println(err)
	// fmt.Println(n)
	// fmt.Println(bs)

	// n, err = file.Read(bs)
	// fmt.Println(err)
	// fmt.Println(n)
	// fmt.Println(bs)

	n := -1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读取文件到了末尾")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
