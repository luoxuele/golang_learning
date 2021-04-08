package main

import "fmt"

//函数式编程
//回调 函数作参数
//闭包 函数作返回值（使用了外层函数的局部变量）

// 局部变量的生命周期会发生改变，正常的局部变量随着函数调用而创建，随着函数的结束而销毁。
// 但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还要继续使用。

func main() {

	fun1 := increment()
	fmt.Println(fun1())
	fmt.Println(fun1())
	fmt.Println(fun1())

	fun2 := increment()
	fmt.Println(fun2())
	fmt.Println(fun2())
	fmt.Println(fun2())

	//内层函数使用了外层函数的局部变量并作为外层函数的返回值时
	//当外层函数执行后，局部变量应该销毁
	//但由于内层函数有使用，所以又不能销毁，就把这个局部变量和返回的匿名函数绑定，放到其他的地方
	//每次执行一个外层函数，就会形成局部变量和返回匿名函数形成的一个闭包

}

func increment() func() int { //外层函数，返回值是一个func()int类型的函数
	//1.外层函数定义了局部变量
	i := 0

	fun := func() int { //内层函数使用了外层函数的局部变量
		i++
		return i
	}

	//3. 返回内层的匿名函数
	return fun

}
