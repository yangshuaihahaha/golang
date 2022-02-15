package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("sayHello......")
	}
}

func test() {
	//这里我们使用defer + recover
	defer func() {
		//来捕获test发出的Panic
		if err := recover(); err != nil {
			fmt.Println("发生错误......")
		}
	}()
	var myMap map[int]string
	myMap[0] = "123123"
}

func main() {
	go test()
	time.Sleep(time.Second)
}
