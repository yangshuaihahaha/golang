package main

import (
	"fmt"
)

func main() {
	//管道的只读只写
	var chan1 chan<- int
	chan1 <- 1
	close(chan1)
	fmt.Println("chan1:", chan1)
	var chan2 <-chan int
	num := <-chan2
	fmt.Println(num)

}
