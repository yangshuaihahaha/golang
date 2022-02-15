package main

import "fmt"

func main() {
	//1,创建一个int类型的channel
	var intChan chan int
	intChan = make(chan int, 10)
	//2,查看地址
	fmt.Printf("intChan的类型是=%v\n", intChan)
	//3,向管道写入数据
	intChan <- 10
	//4,看看管道的长度和容量
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))
	//5,管道跟map不一样，不能一直增加会死锁

	//6,从管道里面取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2:", num2)
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))
	//7,在没有使用协程的情况下，如果所有的数据都已经取出，就会报死锁
	var num3 int
	num3 = <-intChan
	fmt.Println("num3:", num3)
}
