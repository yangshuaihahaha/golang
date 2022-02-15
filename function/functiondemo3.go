package function

import "fmt"

func main() {
	//匿名函数(只能调用一次)
	func() {
		fmt.Println("这是个匿名函数")
	}()

	//多次调用匿名函数
	func8 := func() {
		fmt.Println("我也是个匿名函数")
	}
	func8()

	//定义带参数的匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	//定义带返回值的匿名函数
	res1 := func(a, b int) int {
		return a + b
	}(10, 20) //函数调用了，将函数的执行结果返回给res1
	fmt.Println(res1)

	res2 := func(a, b int) int {
		return a + b
	} //将匿名函数的值赋值给res2
	fmt.Println(res2(3, 9))

}
