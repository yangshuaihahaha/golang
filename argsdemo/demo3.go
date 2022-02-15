package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行所有的参数", len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
	//s := []int{1, 2, 3, 4, 5, 7}

	//Remove(s, 5)
	//Remove(s, 4)
	//for _, v := range s{
	//	fmt.Println("===", v)
	//}
	seq := []string{"a", "b", "c", "d", "e", "f", "g"}
	result := SliceDelete(seq)
	fmt.Println(result)

}

func Remove(s []int, value interface{}) error {
	for i, v := range s {
		if value == v {
			for _, v := range (s)[:i] {
				fmt.Println("<<<", v)
			}
			for _, v := range (s)[i+1:] {
				fmt.Println(">>>", v)
			}
			s = append((s)[:i], (s)[i+1:]...)
			for _, v := range s {
				fmt.Println("---", v)
			}
			return nil
		}
	}
	return nil
}

func SliceDelete(seq []string) (result []string) {
	// 初始化一个新的切片 seq

	// 指定删除位置
	index := 3

	// 输出删除位置之前和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// seq[index+1:]... 表示将后段的整个添加到前段中
	// 将删除前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	// 输出链接后的切片
	return seq
}
