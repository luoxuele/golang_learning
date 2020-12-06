package main

import "fmt"
import "unsafe"


func main() {

	var inter1 int8
	var inter2 int16
	var inter3 int32
	var inter4 int64

	fmt.Printf("%T---%v---%d\n",inter1,inter1,unsafe.Sizeof(inter1))
	fmt.Printf("%T---%v---%d\n",inter2,inter2,unsafe.Sizeof(inter2))
	fmt.Printf("%T---%v---%d\n",inter3,inter3,unsafe.Sizeof(inter3))
	fmt.Printf("%T---%v---%d\n",inter4,inter4,unsafe.Sizeof(inter4))

	var a int
	var b uintptr
	fmt.Printf("%T---%v---%d\n",a,a,unsafe.Sizeof(a))
	fmt.Printf("%T---%v---%d\n",b,b,unsafe.Sizeof(b))

	var c rune
	var d byte
	fmt.Printf("%T---%v---%d\n",c,c,unsafe.Sizeof(c))
	fmt.Printf("%T---%v---%d\n",d,d,unsafe.Sizeof(d))

	
}



// int8	int16	int32	int64
// uint8 uint16	uint32	uint64

//int uint 平台相关

//unicode字符：
	// rune == int32
	// byte == uint8

//uintptr  可以容纳指针