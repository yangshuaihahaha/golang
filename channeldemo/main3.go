package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 1
	intChan <- 2
	close(intChan)
	a := <-intChan
	fmt.Println(a)
}
