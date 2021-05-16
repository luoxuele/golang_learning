package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	fileName := "./ab.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bs := []byte{0}
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(4, io.SeekStart)
	file.Read(bs)
	fmt.Println(string(bs))
	file.WriteString("A")

}

// func (f *File) Seek(offset int64, whence int) (ret int64, err error)
