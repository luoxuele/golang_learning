package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	str := "/home/luoxue/Desktop/cn_office_professional_plus_2019_x86_x64_dvd_5e5be643.iso"
	f, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x\n", h.Sum(nil))
}
