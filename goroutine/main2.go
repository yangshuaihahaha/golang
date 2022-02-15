package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock是一个全局的互斥锁
	lock sync.Mutex
)

//test函数就是计算n!, 让这个结果放入myMap
func test1(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main() {
	//在这里我们开启200个协程完成这些任务
	for i := 1; i <= 20; i++ {
		go test1(i)
	}

	//休眠十秒钟
	time.Sleep(time.Second * 5)

	//这里输出结果，遍历这个结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}
