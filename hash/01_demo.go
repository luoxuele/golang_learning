package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {

	h := md5.New()
	io.WriteString(h, "hello")
	io.WriteString(h, "hehe")

	fmt.Printf("%x\n", h.Sum(nil))

}
