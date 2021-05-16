package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile := "/home/luoxue/Desktop/pdf/HCIA-Datacom V1.0 培训教材.pdf"
	destFile := "hcia.pdf"
	total, err := CopyFile1(srcFile, destFile)
	fmt.Println(total, err)

}

//通过io操作实现文件的拷贝，返回值是拷贝的总数量（字节）和错误
func CopyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}

	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer file1.Close()
	defer file2.Close()

	//读写
	bs := make([]byte, 1024, 1024)
	n := -1 //读取的数据量
	total := 0

	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完成")
			break
		} else if err != nil {
			fmt.Println("报错了")
			return total, err
		}

		total += n
		file2.Write(bs[:n])
	}

	return total, nil
}
