package slice

import "fmt"

func main() {
	var ss []string
	println("func print", ss)
	//切片尾部追加元素append elemnt
	for i := 0; i < 10; i++ {
		ss = append(ss, fmt.Sprintf("s%d", i))
	}
	println("after append", ss)
	for index, value := range ss {
		fmt.Printf("index is:%d,value is %s\n", index, value)
	}
	////删除切片元素remove element at index
	index := 5
	ss = append(ss[:index], ss[index+1:]...)
	println("after delete--------------")

	for index, value := range ss {
		fmt.Printf("index is:%d,value is %s\n", index, value)
	}

	//在切片中间插入元素insert element at index;
	//注意：保存后部剩余元素，必须新建一个临时切片
	rear := append([]string{}, ss[index:]...)
	ss = append(ss[0:index], "inserted")
	ss = append(ss, rear...)
	print("after insert", ss)
}
