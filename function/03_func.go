package main

import "fmt"

func main() {
	sum := getSum(12, 345, 33)
	fmt.Println(sum)

}

//可变参数 ，只能有一个，有多个形参，可变参数要放最后
//可变参数实际上是一个切片
func getSum(nums ...int) int {
	//fmt.Printf("%T %v\n", nums, nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}
