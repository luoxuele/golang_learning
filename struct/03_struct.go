package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      //匿名字段，默认包含了Human的所有字段
	sepciality string
}

func main() {
	mark := Student{Human{"mark", 25, 120}, "Computer Science"}
	fmt.Printf("%T %#v\n", mark, mark)
}
