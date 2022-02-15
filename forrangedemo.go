package main

import "fmt"

func main() {
	arr1 := [...]int{1, 2, 3, 4, 5, 6}
	for index, value := range arr1 {
		fmt.Printf("index的值是:%d, value的值是:%d\n", index, value)
	}
}
