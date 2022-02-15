package main

import "fmt"

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	resChan := make(chan bool, 4)
	//放入数据
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
	//开启四个协程开始工作
	go getPrime(intChan, primeChan, resChan)
	go getPrime(intChan, primeChan, resChan)
	go getPrime(intChan, primeChan, resChan)
	go getPrime(intChan, primeChan, resChan)

	//这里我们主线程进行处理
	go func() {
		for i := 1; i <= 4; i++ {
			<-resChan
		}
	}()

	//遍历所有的素数
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		fmt.Printf("取出的素数是%d\n", res)
	}
	fmt.Println("主线程推出。。。")
}

func getPrime(intChan chan int, primeChan chan int, resChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		//判断是不是素数
		flag := true
		for i := 2; i < num; i++ { //说明不是素数
			if num%i == 0 {
				flag = false
				break
			}
		}
		if !flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个协程因为取不到东西退出了。。。")
	//这里我们还不能关闭primeChan，因为可能其他的管道还在使用。。。
	resChan <- true
}
