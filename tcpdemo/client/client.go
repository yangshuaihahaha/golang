package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Print("client connect error=", err)
		return
	}
	fmt.Println("client connect success conn=", conn)
	//功能一：发送单行数据然后退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入[终端]
	//从终端读取一行数据，并发送给服务器
	line, err := reader.ReadString('\n')
	if err != err {
		fmt.Println("reader.ReadString 失败：", err)
	}
	//再将line发送给服务器
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("conn.Write error:", err)
	}
	fmt.Printf("客户端发送了 %d 字节的数据，并退出了\n", n)
}
