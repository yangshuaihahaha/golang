package function

import "fmt"

func main() {
	/*
		高阶函数：根据go语言的数据类型，可以将一个函数作为另一个函数的参数使用
		fun1, fun2
		将fun1的函数作为fun2的参数
			fun2的函数：就叫高阶函数
				接受了一个函数作为参数，就叫高阶函数
			fun1的函数：
				作为另一个函数的参数，就叫回调函数
	*/
	func4(12, func5)

	//参数直接传入匿名函数
	func4(16, func(a int) int {
		return a
	})
	{

	}
}

func func4(a int, b func(a int) int) {
	fmt.Println(b(a))
}

func func5(a int) int {
	fmt.Println("这是func5的函数")
	return a
}
