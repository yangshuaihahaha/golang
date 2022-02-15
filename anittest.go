package main

import "fmt"

func MyTest() {
	fmt.Println("MyTest函数")
}

func init() {
	fmt.Println("MyTest中的init函数")
}
