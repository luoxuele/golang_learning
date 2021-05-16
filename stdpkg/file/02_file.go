package main

import (
	"fmt"
	"os"
)

// func Mkdir(name string, perm FileMode) error
// func MkdirAll(path string, perm FileMode) error

// func Open(name string) (*File, error)
// func OpenFile(name string, flag int, perm FileMode) (*File, error)
// flag read-only write-only read-write

func main() {
	// err := os.Mkdir("xx/dd", os.ModePerm) //0777
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err := os.MkdirAll("xx/dd/ee", os.ModePerm)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	file, err := os.Open("xxx.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n", file)

}

// const (
// 	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
// 	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
// 	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
// 	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
// 	// The remaining values may be or'ed in to control behavior.
// 	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
// 	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
// 	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
// 	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
// 	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
// )
