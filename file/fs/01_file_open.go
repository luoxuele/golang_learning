package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("./os.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fileObj.Close()

	//read
	var tmp = make([]byte, 128)
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("读了%d个字节byte\n", n)
	fmt.Println(string(tmp[:n]))
}
