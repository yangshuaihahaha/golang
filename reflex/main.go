package main

import (
	"fmt"
	"reflect"
)

func reflexTest(b interface{}) {
	//通过反射获取传入的变量type，kind，的值
	//1，先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	//1，获取到reflect.Value的值
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)
	//注意：rVal并不是数字10，如果使用(rVal + 2)这样是会报错的，获取大真正的value：
	num := rVal.Int()
	fmt.Println("num=", num)
	//下面我们将rVal()转成interface{}
	iV := rVal.Interface()
	num1 := iV.(int)
	fmt.Println("num1=", num1)
}

func reflexTest01(b interface{}) {
	//通过反射获取传入的变量type，kind，的值
	//1，先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	//1，获取到reflect.Value的值
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

	//下面我们将rVal()转成interface{}
	iV := rVal.Interface()
	fmt.Printf("iV = %v, iV = %t\n", iV, iV)
	//将interface通过断言转成需要的类型
	student, ok := iV.(Student)
	if ok {
		fmt.Println("student.name = ", student.Name)
		fmt.Println("student.age = ", student.Age)
	}
	//获取变量的kind
	fmt.Println("rVal.kind = ", rVal.Kind())
	fmt.Println("rType.kind = ", rType.Kind())
	//Type是类型，Kind是类别，Type和Kind可能是相同的，也可能是不同的
	//比如：var num int = 10, num的Type是int，Kind也是int
	//比如：var stu Student stu的Type是pkg1.Student，Kind是struct
}

func main() {
	//演示(基本数据类型，interface，reflex.value)进行反射的基本操作
	//var num int = 10;
	//reflexTest(num)

	//演示结构体的操作
	stu := Student{"Tom", 23}
	reflexTest01(stu)
}

type Student struct {
	Name string
	Age  int
}
