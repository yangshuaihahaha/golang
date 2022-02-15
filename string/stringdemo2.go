package string

import (
	"fmt"
	"strconv"
)

func main() {
	//strconv:字符串和基本数据类型之间的转换
	//boolean
	s1 := "true"
	b1, error := strconv.ParseBool(s1)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(b1)

	ss1 := strconv.FormatBool(b1)
	fmt.Println(ss1)

	//整数
	s2 := "100"
	i2, error := strconv.ParseInt(s2, 10, 64)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Printf("%T,%d\n", i2, i2)

	ss2 := strconv.FormatInt(i2, 10)
	fmt.Println(ss2)

	//itoa(), atio() Atoi是ParseInt(s, 10, 0)的简写。
	i3, error := strconv.Atoi("-42")
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Printf("%T, %d\n", i3, i3)

	ss3 := strconv.Itoa(-42)
	fmt.Printf("%T, %s\n", ss3, ss3)
}
