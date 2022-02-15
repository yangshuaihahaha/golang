package main

import "fmt"

func main() {
	p1 := Person3{"王二狗", 20}
	s1 := Student3{p1, 12, "清华"}
	s1.eat()        //子类对象调用父类方法
	s1.goToSchool() //子类对象调用自己新增方法
	s1.eat()        //如果存在方法的重写, 子类对象访问重写的方法
}

type Student3 struct {
	Person3
	class  int
	school string
}

type Person3 struct {
	name string
	age  int
}

func (p Person3) eat() {
	fmt.Println("人吃饭")
}

func (s Student3) eat() {
	fmt.Println("学生吃饭")
}

func (s Student3) goToSchool() {
	fmt.Println("学生上学")
}
