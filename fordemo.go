package main

import "fmt"

func main() {
	for a := 1; a <= 5; a++ {
		fmt.Println("i的值是", a)
	}

	a := 1
	for a <= 5 {
		fmt.Println(a)
		a++
	}
}
