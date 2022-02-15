package slice

import "fmt"

func main() {
	/*深拷贝：拷贝的是数据的本身
		值类型的数据，默认都是深拷贝：array,int,float,string,bool, struct
	浅拷贝：拷贝的是数据的地址
		导致多个变量指向同一内存
		引用类型的数据，默认都是浅拷贝：slice,map

		因为切片是引用类型的数据，直接拷贝的是地址
	*/
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0)
	for index, value := range s1 {
		fmt.Println(index)
		s2 = append(s2, value)
	}
	fmt.Println(s1, s2)

	//copy()
	s3 := []int{7, 8, 9}
	fmt.Println(s2, s3)

	s2 = s3
	fmt.Println(s2, s3)

	s3[2] = 900
	fmt.Println(s2, s3)

	copy(s3, s2)
	fmt.Println(s2, s3)

	copy(s3[1:], s2[2:])
	fmt.Println(s2, s3)

	s3[2] = 900
	fmt.Println(s2, s3)
}
