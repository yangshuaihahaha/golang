package utils

import "fmt"

func Count() {
	/*
		关于包的使用:
		1,一个目录下的统计文件属于一个包, package声明要一致
		2,package声明的包对应的目录名可以不一致. 但是习惯上还是写一致
		3,包是可以嵌套
		4,同包下的函数不需要导包, 不可以直接使用
		5,main包, main()函数所在的包, 其他的包不能使用
		6,导入包的时候, 路径要从src下面开始
	*/
	fmt.Println("utils包下的Count函数")
}
