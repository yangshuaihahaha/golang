package main

import (
	"fmt"
	"strconv"
)

func main() {
	//使用select解决从管道去数据阻塞的问题
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + strconv.Itoa(i)
	}
	//传统的方法在遍历管道时，如果不关闭就会造成管道阻塞
	//问题是在实际的开发中不好确定什么时候关闭管道
	//可以使用select的方式解决
label:
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan获取的数据是:%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan获取的数据是:%s\n", v)
		default:
			break label
			fmt.Println("都取不到了")
		}
	}
}
