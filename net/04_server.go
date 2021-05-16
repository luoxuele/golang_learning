package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 1. 创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	defer listener.Close()

	//2. 监听客户端连接的请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err: ", err)
			return
		}

		//具体完成服务器和客户端的数据通信
		go HandlerConnet(conn)
	}

}

func HandlerConnet(conn net.Conn) {
	defer conn.Close()

	//获取连接的客户端 addr
	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端成功连接")

	//循环读取客户端发送的数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("服务器检测到客户端已关闭，断开连接！！！")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err: ", err)
			return
		}
		fmt.Println("服务器读到数据：", string(buf[:n]))

		//小写转大写，回发给客户端
		strings.ToUpper(string(buf[:n]))
	}
}
