package main

import "fmt"

func main() {
	fmt.Println(getMul(1, 2, 4, 3.3))
	fmt.Println(getMul2(3.0, 3.2))
}

//add substract multiply divide

func getMul(nums ...float64) float64 {
	mul := 1.0
	for i := 0; i < len(nums); i++ {
		mul *= nums[i]
	}

	return mul
}

//在返回参数列表里面定义了mul,只要return就行了
func getMul2(nums ...float64) (mul float64) {
	mul = 1.0
	for i := 0; i < len(nums); i++ {
		mul *= nums[i]
	}
	return
}
