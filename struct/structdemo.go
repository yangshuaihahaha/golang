package _struct

import "fmt"

func main() {
	/*
		结构体:
			是由一系列具有相同类型或不同类型的数据构成的数据集合
			结构体是由一系列的成员变量构成, 这些成员变量也称为"字段"
	*/
	//创建结构体
	//方法1
	var p1 Person
	fmt.Println(p1)
	p1.name = "王二狗"
	p1.age = 90
	p1.sex = "男"
	p1.address = "北京"
	//方法2
	p2 := Person{}
	p2.name = "王三狗"

	//方法3
	p3 := Person{name: "王大狗", age: 90}
	fmt.Println(p3.name)

	//方法4
	p4 := Person{"王五狗", 90, "男", "甘肃"}
	fmt.Println(p4.age)
}

type Person struct {
	name    string
	age     int
	sex     string
	address string
}
