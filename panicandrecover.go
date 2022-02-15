package main

import "fmt"

func main() {
	/*
		panic: 词义为"恐慌"
		recover: "恢复"
		go语言利用panic(), recover(), 实现程序中特殊的异常处理
			panic(), 让当前的程序进入恐慌, 中断程序的执行
			recover(), 让程序恢复, 必须在defer函数中执行
	*/
	funcA()
	defer myprint("defer main 3......")
	funcB()
	defer myprint("defer main 4......")

	fmt.Println("main over ......")
}

func myprint(s string) {
	fmt.Println(s)
}

func funcA() {
	fmt.Println("函数funcA()......")
}

func funcB() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "程序恢复执行了")
		}
	}()
	fmt.Println("函数funcB()......")
	defer myprint("defer funcB(): 1......")
	for i := 0; i <= 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			panic("funcB, 恐慌了")
		}
	}
	//当外围函数的代码当中发生了运行恐慌, 只有其中所有defer的函数全部执行完毕后, 该运行恐慌才会真正被扩展至调用处。
	defer myprint("defer funcB(): 2......")
}
