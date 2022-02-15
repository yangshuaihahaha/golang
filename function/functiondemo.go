package function

import "fmt"

func main() {
	//go语言至少有一个函数
	/*func funcName (parametername type1, parametername type2) (output1 type1, output2 type2) {
	处理逻辑代码
	返回多个值
	return value1,value2
	*/
	/*
		func : 函数的声明由func开始
		funcName: 函数的名称
		()：函数的标志
		参数列表：形式参数用于接受外部传入函数的数据
		返回值列表：函数执行后返回给调用处的结果
	*/
	getSum(20)

	//多个参数
	getAdd(10, 10)

	//类型一致的话可以省略
	getAdd2(10, 10, 10)
	getAdd3("就", "是", 10)

	/*
		可变参数：一个函数的类型确定，个数不确定
		传入的参数就是一个切片
		注意:
			1,如果一个函数的参数是可变参数，同时还有其他的参数，可变参数要放在参数列表的最后
			2,一个函数最多只有一个可变参数
	*/
	getSum1(1, 2, 3, 4, 5, 6, 7, 8, 9)
	s1 := []int{1, 2, 3, 4, 5, 6}
	getSum1(s1...)

	fmt.Println("值传递----------------------------------------------")
	/*
		值传递和引用传递
			A: 值传递:传递的是数据的本身。修改数据对于原始数据没有影响。
				值类型的数据, 默认都是值类型：基础类型，array, struct
			B: 引用类型：传递的是数据的地址，导致多个变量指向同一块内存。
				引用类型的数据，默认都是引用传递：slice， map，chan
	*/
	var a int = 10
	fun1(a)
	fmt.Println(a)

	/*
		函数的返回值
	*/
	//定义一个函数，带返回值
	fmt.Println(fun2(1))
	//定义函数时指明返回数据是哪一个
	fmt.Println(fun3(12))
	//函数的多个返回值
	a, b := fun4(12, 13)
	fmt.Println("周长是", a)
	fmt.Println("面积是", b)
	c, d := fun6(2, 3)
	fmt.Println("finc6函数的周长是", c)
	fmt.Println("finc6函数的面积是", d)
}

func getSum(n int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func getAdd(a int, b int) {
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}

func getAdd2(a, b, c int) {
	sum := a + b + c
	fmt.Printf("%d + %d + %d = %d\n", a, b, c, sum)
}

func getAdd3(a, b string, c int) {
	fmt.Printf("a:%s b:%s c:%d\n", a, b, c)
}

func getSum1(num ...int) {
	sum := 0
	for value := range num {
		sum += value
	}
	fmt.Println("总和是", sum)
}

func fun1(a int) {
	a = 20
}

func fun2(a int) int {
	return 20
}

func fun3(a int) (sum int) {
	sum = 3
	return
}

func fun4(a int, b int) (int, int) {
	perimeter := (a + b) * 2
	area := a * b
	return perimeter, area
}

func fun6(a int, b int) (perimeter int, area int) {
	perimeter = (a + b) * 2
	area = a * b
	return
}
