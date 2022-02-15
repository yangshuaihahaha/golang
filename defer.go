package main

import "fmt"

func main() {
	/*
		defer的词义："延迟，推迟"
		在go语言中，使用defer关键字来延迟一个函数或者方法的执行。

		1,defer函数或者方法：一个函数或方法的执行被延迟了
		2,defer的用法
			A:对象.close(),临时文件的删除。。。
				文件.open()
				defer close()
				读或写
			B: go语言中关于异常的处理, 使用panic()和recover()
				panic函数用于引发恐慌，导致程序中断执行
				recover函数用于恢复程序的执行，recover()语法上要求必须在defer中执行
		3,如果有多个defer函数:

		4,defer函数传递参数的时候:
			defer在函数调用的时候就已经发生参数传递了,只不过是没有执行而已
			{
				a := 2
				fmt.Println(a)
				defer func2(a)
				a++
				fmt.Println("main 函数中的a值是", a)
			}

		5,defer函数的注意点:
			当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行
			当执行外围函数的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回
			当外围函数的代码引发运行恐慌时，只有其中所有的延迟函数都执行完毕后，改运行时恐慌才会真正扩展至调用函数

	*/
	//defer func1("hello")
	//defer fmt.Println("123456")
	////如果在一个函数中使用defer调用了另一个函数，此函数就是外围函数，而被defer调用的函数会在最后执行
	//defer func1("world")
	//fmt.Println("王二狗")

	fmt.Println(func3())
}

func func3() int {
	fmt.Println("func3 函数的执行")
	defer func1("haha")
	return 0
}

func func2(s1 int) {
	fmt.Println("func2 函数中a的值是", s1)
}

func func1(s1 string) {
	fmt.Println(s1)
}
