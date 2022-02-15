package main

import "fmt"

func main() {
	/*
		方法: method
			一个方法就是一个包含了接受者的函数, 接受者可以是命名类型或者结构体类型的一个值或者一个指针
			所有给定类型的方法属于类型的方法集

		对比函数:
			A: 意义
				方法: 某个类别的行为功能, 需要指定的接受者调用
				函数: 一段独立功能的代码, 可以直接调用
			B: 语法
				方法: 方法名可以相同, 只要接受者不同
				函数: 命名不能冲突

	*/
	//结构体类型的一个值调用
	w1 := Worker1{"王二狗", 20, "男"}
	w1.work()
	//结构体一个指针调用
	w2 := &Worker1{"Tom", 21, "男"}
	w2.work()

	w1.rest()
	w2.rest()

}

//1,定义一个工人体结构
type Worker1 struct {
	name string
	age  int
	sex  string
}

//2,定义方法行为
func (w Worker1) work() {
	fmt.Println(w.name, "在工作")
}

func (p *Worker1) rest() {
	fmt.Println(p.name, "在休息")
}
