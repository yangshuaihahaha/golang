package _interface

import "fmt"

func main() {
	/*
		接口的嵌套
	*/

	c1 := Cat1{}
	c1.test1()
	c1.test2()
	c1.test3()

	var a1 A1 = c1
	a1.test1()

	var b1 B1 = c1
	b1.test2()

	var cc1 C1 = c1
	cc1.test1()
	cc1.test2()
	cc1.test3()

}

type A1 interface {
	test1()
}

type B1 interface {
	test2()
}

type C1 interface {
	A1
	B1
	test3()
}

type Cat1 struct {
}

func (c Cat1) test1() {
	fmt.Println("test1()")
}
func (c Cat1) test2() {
	fmt.Println("test2()")
}
func (c Cat1) test3() {
	fmt.Println("test3()")
}
