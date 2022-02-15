package _interface

import (
	"fmt"
	"math"
)

func main() {
	/*
		接口的断言
	*/

	c1 := Circle{2}
	fmt.Println(c1.prei())

	t1 := Triangle{3, 4, 5}
	fmt.Println(t1.prei())

	var s1 Shape
	s1 = c1
	testShape(s1)
	s1 = t1
	testShape(s1)

}

func testShape(s Shape) {
	//断言(其中发生的是值传递, 不是引用类型的数据)

	//方法一
	if ins, ok := s.(Triangle); ok {
		fmt.Println("三角形其中一条边长是:", ins.a)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("圆形的半径是:", ins.radius)
	}

	//方法二
	switch ins := s.(type) {
	case Circle:
		fmt.Println("圆形的半径是:", ins.radius)
	case Triangle:
		fmt.Println("三角形其中一条边长是:", ins.a)
	}
}

//1, 定义一个接口
type Shape interface {
	prei() float64
}

//2,定义实现类三角形
type Triangle struct {
	a, b, c float64
}

//3,定义实现类三圆形
type Circle struct {
	radius float64
}

func (t Triangle) prei() float64 {
	return t.a + t.b + t.c
}

func (c Circle) prei() float64 {
	return c.radius * 2 * math.Pi
}
