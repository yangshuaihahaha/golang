package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 3)
	allChan <- Cat{"Tom", 12}
	allChan <- 12
	allChan <- "momo"

	newCat := <-allChan
	fmt.Printf("从管道中取出的newCat=%T, newCat=%v\n", newCat, newCat)
	a := newCat.(Cat)
	fmt.Println(a.Name)
}
