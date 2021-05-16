package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}

	defer conn.Close()

	//获取客户键盘输入（stdin）,将输入数据发送给服务器
	go func() {
		str := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Read err: ", err)
				continue
			}

			//写给服务器，读多少，写多少
			conn.Write(str[:n])
		}
	}()

	//回显服务器发回的大写数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("检查到服务器关闭，客户端也关闭")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err: ", err)
			return
		}
		fmt.Println("客户端读到服务器回发： ", string(buf[:n-1]))
		// fmt.Printf("%0x", string(buf[:n]))
	}
}
