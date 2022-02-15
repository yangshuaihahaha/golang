package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听。。。")
	//tcp表示使用的网络协议是tcp
	listen, error := net.Listen("tcp", "0.0.0.0:8888")
	if error != nil {
		fmt.Println("listen error=", error)
	}
	defer listen.Close()
	//循环等待客户端连接
	for {
		fmt.Println("等待客户端连接。。。。")
		conn, error := listen.Accept()
		if error != nil {
			fmt.Println("accept error=", error)
			return
		}
		fmt.Printf("accept success and conn=%v\n", conn)
		fmt.Println("客户端的ip是", conn.RemoteAddr().String())
		//这里启动一个协程
		go process(conn)
	}
}

func process(conn net.Conn) {
	//这里循环的接受客户端发送的数据
	defer conn.Close() //关闭conn
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//在这里等待conn发送信息，如果客户端没有write，那么就阻塞在这里
		fmt.Println("服务器在等待你的输入。。。")
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器的read err=", err)
			return
		}
		//显示客户端发送的内容到服务器终端
		fmt.Print(string(buf[:n]))
	}
}
