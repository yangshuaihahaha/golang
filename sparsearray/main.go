package main

import "fmt"

//只记录几行，几列，值是什么
type node struct {
	row int
	col int
	val int
}

func main() {
	var charMap [11][11]int
	charMap[1][2] = 1
	charMap[2][3] = 2
	for _, v1 := range charMap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
