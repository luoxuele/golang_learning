package main

import (
	"errors"
	"fmt"
)

var isActive bool
var enabled, disabled = true, false

func main() {

	//1. boolean
	fmt.Printf("%T %v\n", isActive, isActive)
	fmt.Printf("%T %v\n", enabled, enabled)
	fmt.Printf("%T %v\n", disabled, disabled)

	//2. number
	// int(默认) unit	//默认是32位
	//int8(rune) int16 int32 int64
	//uint8(byte) uint16 uint32 uint64
	//float32 float64(默认)
	//complex128 (64位实数+64位虚数)  complex64

	var a = 100
	var b int = 100
	c := a + b //(mismatched types int and int16)
	fmt.Printf("%T %v\n", c, c)

	// 3. 字符串
	no, yes, hello := "no", "yes", "hello world"
	fmt.Printf("%T %v\n", no, no)
	fmt.Printf("%T %v\n", yes, yes)
	fmt.Printf("%T %v\n", hello, hello)

	// 重新赋值
	no = "yes"
	fmt.Printf("%T %v %b\n", no[0], no[0], no[0]) //字符串的单个字符是byte(uint8)

	// no[0] = 'c' //字符串是不可变的
	//但可以通过[]byte转换  (切片)
	temp := []byte(no)
	temp[0] = 'n'
	no = string(temp)
	fmt.Printf("%T %v \n", no, no)
	fmt.Printf("%T %v \n", no+yes, no+yes)

	s := "hello"
	s = "c" + s[1:]
	fmt.Println(s)

	//反引号可以声明一个多行的字符串
	m := `hello
	
			world`

	fmt.Println(m)

	// 4. 错误类型
	err := errors.New("emit macho dwarf: elf header corrupted\n")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%T %v \n", err, err)

	//5 . 分组声明
	var (
		i      = 100
		Pi     = 3.14
		prefix = "Go_"
	)
	fmt.Println(i, Pi, prefix)

	// 6. iota
	const (
		x = iota //0
		y = iota //1
		z
		w
	)
	fmt.Println(x, y, z, w)

	// 7 .golang的一些规则
	// 大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
	// 大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。

}
