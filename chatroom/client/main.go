package main

import (
	"fmt"
	"os"
)

//定义两个变量，一个表示用户id，一个表示用户名密码
var userId int
var userPwd string

func main() {
	//用于接受用户的选择
	var key int
	//判断是否继续显示菜单
	var loop = true

	for loop {
		fmt.Println("---------欢迎登陆多人聊天系统---------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 1 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("您的输入有误，请重新输入")
			loop = false
		}
	}

	//根据用户的输入，显示新的信息
	if key == 1 {
		//说明用户要登陆
		fmt.Println("请输入用户的id")
		fmt.Scanln(&userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanln(&userPwd)
		//先把登陆的函数写到另一个文件，比如login.go
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("用户登录失败，用户名或者密码有误")
		} else {
			fmt.Println("登陆成功")
		}
	} else if key == 2 {
		fmt.Println("进行用户注册的逻辑。。。")
	} else if key == 3 {
		os.Exit(0)
	}
}
