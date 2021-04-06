package main

import (
	"fmt"
	"unsafe"
)

//在Go中，接口是一组方法签名
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokia NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("I am Iphone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()
	phone = new(Iphone)
	phone.call()

	fmt.Printf("%T %v\n", phone, phone)
	fmt.Println(unsafe.Sizeof(phone))
}

// interface可以被任意的对象实现
// 一个对象可以实现任意多个interface
// 任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface
