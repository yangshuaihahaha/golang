package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 6; i++ {
		if i == 3 {
			goto THREE
		}
		fmt.Println(i)
	}
THREE:
	{
		fmt.Println("这是三")
	}

}
