package main

import "fmt"

func main() {
	var arr1 [4]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1)

	var arr2 = [4]int{1, 2, 3}
	fmt.Println(arr2)

	var arr3 = [6]int{1: 3, 3: 36}
	fmt.Println(arr3)

	var arr4 = [3]string{1: "12", 0: "13", 2: "14"}
	fmt.Println(arr4)

	var arr5 [6]string
	fmt.Println(arr5)

	arr6 := [...]int{1, 2, 4}
	fmt.Println(arr6)
}
