package pointer

import "fmt"

func main() {
	/*
		指针: pointer
			存储了另一个变量的内存地址
	*/
	//1. 定义一个int类型的变量
	a := 10
	fmt.Println("a的地址是:", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("a的地址是: %p\n", &a)

	//2,创建一个指针变量, 用于存储变量a的地址
	var p1 *int
	fmt.Println(p1)
	p1 = &a
	fmt.Println("p1的数值:", p1) //p1中存储的是a的地址
	fmt.Printf("p1自己的地址是:%p\n", &p1)
	fmt.Println("p1的数值，是a的地址，该地址存储的数据：", *p1) //获取指针指向的变量的树枝
}
