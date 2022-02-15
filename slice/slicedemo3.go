package slice

import "fmt"

func main() {
	/*按照类型来分：
		基本类型：int,float.string,bool
		复合类型：array,slice,map,struct,pointer,function,chan
	按照特点来分：
		值类型：int,float,string,bool,array
	   		传递的是数据的副本
		引用类型：slice
	 		传递的是地址，多个变量志向同一个地址

	所以：切片是引用类型的数据，储存了底层数组的引用
	*/
	a1 := [4]int{1, 2, 3, 4}
	a2 := a1
	fmt.Println(a1, a2)
	a2[2] = 100
	fmt.Println(a1, a2)

	s1 := []int{1, 2, 3, 4}
	s2 := s1
	fmt.Println(s1, s2)
	s2[2] = 100
	fmt.Println(s1, s2)
}
