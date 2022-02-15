package main

import "fmt"

func main() {
	/*
		面向对象: OOP
	*/
	//1.创建父类结构
	p1 := Person2{"Tom", 12}

	//2.创建子类对象
	s1 := Student2{p1, 12}
	fmt.Println(s1.Person2.age)

	/*提升字段
	Student2结构体将Person2结构体作为一个匿名字段了
	那么Person2中的字段,对于Student2来讲, 就是提升字段
	Student2直接能访问Person2中的字段
	*/
	fmt.Println(p1.age) //等于s1.Person2.age
}

//1.定义父类
type Person2 struct {
	name string
	age  int
}

//2.定义子类
type Student2 struct {
	Person2  //模拟继承结构
	classnum int
}
