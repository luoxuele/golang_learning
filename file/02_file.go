package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	// 路径：
	// 	相对路径 relative
	// 	绝对路径 /

	fileName1 := "/home/luoxue/git/golang_learning/file/xxx.txt"
	fileName2 := "./hhh.txt"

	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))

	fmt.Println(path.Join(fileName1, ".."))

	//创建目录
	// err := os.Mkdir("/home/luoxue/git/golang_learning/file/dd", os.ModePerm)
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }
	// fmt.Println("目录创建成功")

	// err := os.MkdirAll("./cc/c1/c2/c3", os.ModePerm)
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }
	// fmt.Println("创建成功")

	//创建文件
	// file1, err := os.Create("./cre.txt")
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }

	// fmt.Println(file1)

	//打开文件
	// file3, err := os.Open(fileName1) //readonly
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// 	return
	// }
	// fmt.Println(file3)

	file4, err := os.OpenFile(fileName1, os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println(file4)

}
