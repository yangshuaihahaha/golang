package _struct

import "fmt"

func main() {
	/*
		结构体的嵌套: 一个结构体中的字段,是另一个结构体

	*/

	b := Book{Stu{90}, 12}
	fmt.Println(b.student.age)
}

type Book struct {
	student Stu
	age     int
}

type Stu struct {
	age int
}
