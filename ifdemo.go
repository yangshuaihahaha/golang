package main

import "fmt"

func main() {
	num := 11
	if num > 10 {
		fmt.Println("是比10大")
	}

	if num2 := 6; num2 < 3 {
		fmt.Println("进入循环")
	} else if num2 > 4 {
		fmt.Println("第二次进入循环")
	}
}
