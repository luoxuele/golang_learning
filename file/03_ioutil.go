package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// byteStr, err := ioutil.ReadFile("/home/luoxue/Desktop/cn_office_professional_plus_2019_x86_x64_dvd_5e5be643.iso")
	byteStr, err := ioutil.ReadFile("/media/luoxue/系统1/baidudownload/CentOS-7.4-x86_64-DVD-1708.iso")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(string(byteStr))
	fmt.Printf("%v   %v\n", cap(byteStr), len(byteStr))
}
