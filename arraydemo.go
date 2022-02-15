package main

import (
	"demo/utils"
	"demo/utils/person"
	"fmt"
)

func main() {
	arr := [3][4]int{{1, 2, 3}, {1, 2, 3, 4}, {1, 2, 3, 4}}
	var arr2 [3][4]int
	fmt.Println(arr)
	fmt.Println(arr2)
	utils.Count()
	p1 := person.Person{"123"}
	fmt.Println(p1.Name)
}
