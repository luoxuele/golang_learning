package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "./ab.txt"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	defer file.Close()

	//写数据
	// bs := []byte{65, 66, 67, 68, 69, 70} //ABCDEF
	bs := []byte{97, 98, 99, 100, 101, 102}
	n, err := file.Write(bs[2:])
	fmt.Println(n)
	HandleErr(err)

	n, err = file.WriteString("hello world")
	fmt.Println(n)
	HandleErr(err)

}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
