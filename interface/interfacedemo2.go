package _interface

import (
	main2 "demo"
	"fmt"
)

func main() {
	/*
		空接口(interface{})
			不包含任何的方法, 正因为如此, 所有的类型都实现了接口, 因为空接口可以储存任意类型的数值

	*/
	var a1 A = Cat{"花猫"}
	var a2 A = Person4{"王二狗", 12}
	var a3 A = "haha"
	var a4 A = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	test2(a1)
	test2(a2)
	test2(a3)
	test2(a4)

	//map, key字符串, value任意类型
	map1 := make(map[string]interface{})
	map1["name"] = "王二狗"
	map1["age"] = 12
	map1["friend"] = main2.Person3{"jerry", 12}

	//切片,存储任意类型的数据
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, 12)
	slice1 = append(slice1, "hahaha")
	slice1 = append(slice1, main2.Person3{"jerry", 12})

}

//接口A是空接口，理解为代表了任意类型
func test2(a A) {

}
func test1(a interface{}) {

}

//空接口
type A interface {
}

type Cat struct {
	color string
}

type Person4 struct {
	name string
	age  int
}
