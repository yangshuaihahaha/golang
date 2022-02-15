package main

import "fmt"

func main() {
	arr := [6]int{3, 12, 9, 90, 2, 35}
	for index, value := range arr {
		fmt.Printf("index的值是:%d, value的值是:%d\n", index, value)
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}
