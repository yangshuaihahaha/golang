package main

import (
	message2 "demo/chatroom/common/message"
	"fmt"
	"net"
)

func login(userId int, userPwd string) (error error) {
	//fmt.Printf("用户输入的userId是: %d, 用户输入的密码是: %s\n", userId, userPwd)
	//1，连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}
	//2，准备通过conn发送消息给服务
	var message message2.Message
	message.Type = message2.LoginMesType
	var loginMes message2.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	return nil
}
