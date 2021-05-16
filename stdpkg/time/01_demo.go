package main

import (
	"fmt"
	"time"
)

func main() {

	// 1. 获取当前时间
	t1 := time.Now()
	fmt.Printf("%T\n", t1)
	fmt.Println(t1)

	// 2. 获取指定的时间
	t2 := time.Date(2008, 7, 15, 16, 30, 28, 0, time.Local)
	fmt.Println(t2)

	//time->string
	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println("s1:\t", s1)

	s2 := t1.Format("2006/01/02")
	fmt.Println("s2:\t", s2)

	s3 := "1994年07月15日"
	t3, err := time.Parse("2006年01月02日", s3)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(t3)
	fmt.Printf("%T\n", t3)

}
