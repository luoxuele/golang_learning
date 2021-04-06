package main

import "fmt"

func main() {

	// fmt.Println(rectangle(3, 4))
	//不要的返回值可以有“_”丢弃
	//"_" : 空白标识符，用来舍弃数据
	_, area := rectangle(4, 8)
	fmt.Println(area)

	fmt.Println(rectangle2(5, 5))

}

func rectangle(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area
}

func rectangle2(len, wid float64) (perimeter, area float64) {
	perimeter = (len + wid) * 2
	area = len * wid
	return
}
