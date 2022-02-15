package _struct

import "fmt"

func main() {
	/*
		匿名结构体和匿名字段:

		匿名结构体: 没有名字的结构体
			在创建匿名结构体的时候, 同时创建对象
		匿名字段: 一个结构体的字段没有字段名
	*/
	//正常结构体
	s1 := Student{name: "zhangsan"}
	fmt.Println(s1.name)

	//匿名结构体
	s2 := struct {
		name string
		age  int
	}{"李四", 12}
	fmt.Println(s2.age)

	//匿名字段(默认把数据类型作为名字)
	w2 := Woker{"李小华", 32}
	fmt.Println(w2.string)
}

type Student struct {
	name string
	age  int
}

type Woker struct {
	string
	int
	//string(字段不能重复)
}
