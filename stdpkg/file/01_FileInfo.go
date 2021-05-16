package main

import (
	"fmt"
	"os"
)

func main() {
	// fileInfo, err := os.Stat("./01_FileInfo.go")
	fileInfo, err := os.Stat(".")

	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Printf("%T\n", fileInfo) //*os.fileStat
	fmt.Printf("%+v\n", fileInfo)

	xx := fmt.Println
	// xx(fileInfo.Name())
	// xx(fileInfo.Size())
	// xx(fileInfo.Mode())
	xx(fileInfo.ModTime())
	xx(fileInfo.IsDir())
	fmt.Printf("%+v\n", fileInfo.Sys())

}

// // A FileInfo describes a file and is returned by Stat.
// type FileInfo interface {
// 	Name() string       // base name of the file
// 	Size() int64        // length in bytes for regular files; system-dependent for others
// 	Mode() FileMode     // file mode bits
// 	ModTime() time.Time // modification time
// 	IsDir() bool        // abbreviation for Mode().IsDir()
// 	Sys() interface{}   // underlying data source (can return nil)
// }

// A fileStat is the implementation of FileInfo returned by Stat and Lstat.
// type fileStat struct {
// 	name    string
// 	size    int64
// 	mode    FileMode
// 	modTime time.Time
// 	sys     syscall.Stat_t
// }
