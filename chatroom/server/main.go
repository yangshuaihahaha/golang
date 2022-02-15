package server

import (
	"fmt"
	"net"
)

//处理和客户端的通讯
func process(conn net.Conn) {
	//读取客户端发送的信息
}

func main() {
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("net.Listen error:", err)
	}
	//一旦连接成功，就等待客户端进行连接
	for {
		fmt.Println("等待客户端来进行连接。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept error:", err)
		}
		//连接成功，则启动一个协程保持客户端通讯
		go process(conn)
	}
}
