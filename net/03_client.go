package main

import (
	"fmt"
	"net"
)

func main() {
	//指定服务器 ip+port 创建通信套接字
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}

	defer conn.Close()

	//主动写数据给服务器
	conn.Write([]byte("Are you Ready?"))

	//接受服务器回发的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}

	fmt.Println("服务器回发：", string(buf[:n]))

}
