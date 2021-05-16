package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	defer listener.Close()

	//阻塞等待用户连接

}
