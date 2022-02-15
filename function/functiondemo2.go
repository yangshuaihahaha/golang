package function

import "fmt"

func main() {
	//1.函数就是一个变量
	fmt.Printf("%T\n", fun7) //func(int, int)
	fmt.Println(fun7)        //0x109af30 函数名对应函数体的具体地址

	//2.直接定义一个函数类型的变量
	var c func(int, int)
	fmt.Println(c)

	c = fun7 //将fun1的值(函数的具体地址)赋值给c
	fmt.Println(c)

	c(10, 0) //c也是函数类型，加小括号也可以被调用

}

func fun7(a, b int) {
	fmt.Println("函数func7被调用")
}
