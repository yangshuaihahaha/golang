package main

import "fmt"

func main() {
	num := 5
	switch num {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	default:
		fmt.Println("数据有错误")
	}

	switch {
	case 1 > 2:
		fmt.Println("1大于2")
	case 1 < 2:
		fmt.Println("1小于2")
	}

	switch a := 1; a {
	case 1, 2, 3:
		fmt.Println("是1，2，3中的数值")
	case 4, 5, 6:
		fmt.Println("是4，5，6中的数值")
	}

	switch b := 1; b {
	case 1:
		fmt.Println("fallthrough---1")
		fallthrough
	case 2:
		fmt.Println("fallthrough穿透2")
	}
}
