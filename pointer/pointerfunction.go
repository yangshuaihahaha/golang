package pointer

import "fmt"

func main() {
	/*
		函数指针: 一个指针，指向了一个函数的指针
			因为go语言中, function, 默认是一个指针, 没有*
			slice,map,function
		指针函数: 一个函数，该函数的返回值是指针
	*/
	var a = fun8()
	fmt.Printf("%T, %p, %v\n", a, &a, a)

	var b = fun9()
	fmt.Printf("%T, %p, %v\n", b, &b, b)
	fmt.Println(b[2])

}

func fun8() [4]int {
	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("fun8数组的地址%p\n", &arr)
	return arr
}

func fun9() *[4]int {
	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("fun9数组的地址%p\n", &arr)
	return &arr
}
