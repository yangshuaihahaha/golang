package pointer

import "fmt"

func main() {
	/*
		指针作为参数:
			参数的传递: 值传递, 引用传递
	*/
	num := 10
	fun10(&num)
	fmt.Println("fun10()函数运行后的num:", num)
}

func fun10(p1 *int) {
	fmt.Println("fun10()函数中p1:", *p1)
	*p1 = 100
	fmt.Println("fun10()函数中修改的p1:", *p1)
}
