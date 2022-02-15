package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 10
	reflect01(&num)
	fmt.Println(num)
}

//通过反射修改int的值
func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	//.Elem()获取指针的真实地址，而且要想更改其中的值，必须传入地址然后通过Elem()去修改
	rVal.Elem().SetInt(12)
}
