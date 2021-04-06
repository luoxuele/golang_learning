package main

import "fmt"

//方法可以继承和重写

type Human struct {
	name  string
	age   int
	phone string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s ,You can call me on %s\n", h.name, h.phone)
}

type Student struct {
	Human
	school string
}
type Employee struct {
	Human
	company string
}

func (s *Student) SayHi() {
	fmt.Printf("I'm %s, 我在 %s 上学\n", s.name, s.school)
}

func (e *Employee) SayHi() {
	fmt.Printf("I’m %s, 我在 %s 上班\n", e.name, e.company)
}

func main() {
	h1 := Human{"田昌", 27, "16675576315"}
	h1.SayHi()

	zhangsan := Student{Human{"张三", 22, "16658465894"}, "吉首一中"}
	zhangsan.SayHi()

	lisi := Employee{Human{"李四", 31, "13756842596"}, "浙江雪狐科技有限公司"}
	lisi.SayHi()
}
