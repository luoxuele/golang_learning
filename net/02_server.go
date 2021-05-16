package main

import (
	"fmt"
	"net"
)

func main() {
	//创建一个用于监听的sokcet
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()
	//fmt.Println(listener)

	//阻塞监听客户端连接请求，成功建立连接，返回用于通信的socket
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}

	conn.Write(buf[:n]) //读多少写多少
	fmt.Println(string(buf[:n]))

}
