package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	for i := 1; i < 10; i++ {
		fmt.Println("test", 1)
		time.Sleep(time.Second)
	}
}

func main() {
	go test()
	for i := 1; i < 10; i++ {
		fmt.Println("main", 1)
		time.Sleep(time.Second)
	}
	//获取cpu个数
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)
	//自己设置使用的cpu个数
	runtime.GOMAXPROCS(cpuNum)
}
