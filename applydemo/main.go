package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go writeDate(intChan)
	go readDate(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

func writeDate(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("writeDate 写入数据=%v\n", i)

	}
	close(intChan)
}

func readDate(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	exitChan <- true
	close(exitChan)
}
