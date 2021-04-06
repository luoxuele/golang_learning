package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//打开文件
	file, err := os.Open("./01_Open.go")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	//操作文件
	// fmt.Printf("%T\n", *file) //os.File 结构体
	var strSlice []byte
	var tempSlice = make([]byte, 32)

	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}

		if err != nil {
			fmt.Println("读取失败")
			return
		}
		// fmt.Printf("读取到了%v个字节\n", n)
		strSlice = append(strSlice, tempSlice[:n]...)

	}
	fmt.Println(string(strSlice))
	fmt.Printf("size: %v bytes\n", len(strSlice))

}
