package main

import "fmt"

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

func main() {
	var p1 Usber
	p := Phone{Name: "华为手机"}
	p1 = p
	p.start()

	fmt.Printf("%T\n", p1)
}
