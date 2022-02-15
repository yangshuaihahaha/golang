package _type

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		type: 用于类型定义和类型别名
			1, 类型定义: type 类型名 Type
			2, 类型别名: type 类型名 = Type
	*/
	var name mystr
	name = "王二狗"
	var s1 string
	s1 = "李小花"
	fmt.Println(name, s1)
	//name = s1(这里会报错, 因为在语法上完全是两种类型)

	res1 := fun11()
	res1(10, 20)

	//类型别名和类型通用
	var i3 myint2
	i2 := 100
	i3 = i2
	fmt.Println(i3)
}

//1,定义一个新的类型
type mystr string

//2,定义函数类型
type myfun func(int, int) string

func fun11() myfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}

//3, 类型别名
type myint2 = int //不是重新定义新的数据类型, 只是给int起别名, 和int可以通用
