package main

import "fmt"

func main() {
	/*创建map的语法
	 */
	var map1 map[int]string //只有声明没有初始化，是一个nil的map
	var map2 = make(map[int]string)
	var map3 = map[string]int{"Go": 98, "java": 89, "python": 89}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	/*
		每种数据类型都有对应的默认值
			int: 0
			float:0.0 -> 0
			string:""
			array:[00000]

			slice:nil
			map:nil
	*/
	var a1 [0]int
	fmt.Println(a1)

	//var s1 []int
	//s1[0] = 23
	//fmt.Println(s1)

	var s2 = make([]int, 0)
	fmt.Println(s2)

	//判断map是不是nil
	if map1 != nil {
		map1 = make(map[int]string)
		fmt.Println(map1)
	}

	//获取map
	fmt.Println(map3["Go"])
	value, ok := map3["Go1"]
	fmt.Println(ok, value)

	//删除数据
	delete(map3, "java")
	fmt.Println(map3["java"])

	//获取长度
	fmt.Println(len(map3))

}
