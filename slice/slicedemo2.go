package slice

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := arr[:5]  //1-5
	s2 := arr[3:8] //4-8
	s3 := arr[5:]  //6-10
	s4 := arr[:]   //1-10
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Printf("长度是%d，容量是%d\n", len(s1), cap(s1))
	fmt.Printf("长度是%d，容量是%d\n", len(s2), cap(s2))
	fmt.Printf("长度是%d，容量是%d\n", len(s3), cap(s3))
	fmt.Printf("长度是%d，容量是%d\n", len(s4), cap(s4))

	fmt.Println("---------更改数组的内容---------")
	arr[4] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("---------更改切片的内容---------")
	s1[2] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("---------更改切片的内容---------")
	s1 = append(s1, 1, 1)
	fmt.Println(arr)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	//超过原来切片a的长度就会创建新的切片，而不是更改原来的切片
	fmt.Println("---------更改切片的内容---------")
	s1 = append(s1, 2, 2, 2, 2)
	fmt.Println(arr)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

}
