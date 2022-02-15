package string

import (
	"fmt"
	"strings"
)

func main() {
	//是否包含某一个字符串
	s1 := "helloword"
	fmt.Println(strings.Contains(s1, "hello"))
	//是否包含char中任意一个字符串
	fmt.Println(strings.ContainsAny(s1, "top"))
	//统计char在string出现的次数
	fmt.Println(strings.Count(s1, "lo"))

	//查找在string中的位置，没有就返回-1
	fmt.Println(strings.Index(s1, "lo"))
	//查找char的任意一个字符在字符串中出现的位置
	fmt.Println(strings.IndexAny(s1, "ello"))
	//查找字符串中最后一次出现的位置
	fmt.Println(strings.LastIndex(s1, "o"))

	//字符串的拼接
	s2 := []string{"qw", "er", "tt", "yu"}
	s3 := strings.Join(s2, "-")
	fmt.Println(s3)

	//字符串的切割
	s4 := "qw,er,gb,cd"
	s5 := strings.Split(s4, ",")
	fmt.Println(s5)

	//重复拼接自己
	s6 := strings.Repeat("hello", 5)
	fmt.Println(s6)

	//替换
	fmt.Println(strings.Replace("hello", "l", "%", 1)) //只替换一个
	fmt.Println(strings.Replace("hello", "l", "%", -1))

	//大小写转换
	fmt.Println(strings.ToLower("HELLo"))
	fmt.Println(strings.ToUpper("HELLo"))

	//字符串的截取,没有substring的方法
	s7 := "Hello, World"
	fmt.Println(s7[:5])
	fmt.Println(s7[5:])
}
