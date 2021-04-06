package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	//1.实例化
	// var p1 Person //对象的首地址
	// p1.Name = "田玲利"
	// p1.Age = 25
	// p1.Sex = "女"
	// fmt.Printf("%T %#v\n", p1, p1)

	// var p2 = new(Person) //指针，指向一个对象
	// p2.Name = "田昌"
	// p2.Age = 27
	// p2.Sex = "男"

	// fmt.Printf("%T %#v\n", p2, p2)

	// var p3 = &Person{}
	// // p3.Name = "甜甜"
	// // p3.Age = 1
	// // p3.Sex = "女"
	// fmt.Printf("%T %#v\n", p3, p3)

	// var p1 Person = Person{"田玲利", 25, "女"}
	p1 := Person{"田玲利", 25, "女"}
	fmt.Printf("%T %#v\n", p1, p1)

	p2 := Person{Age: 27, Name: "田昌"}
	fmt.Printf("%T %#v\n", p2, p2)

	p3 := new(Person)
	p3.Age = 23
	fmt.Printf("%T %#v\n", *p3, *p3)

}
