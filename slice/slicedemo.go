package slice

import "fmt"

func main() {
	var arr = []int{1, 2, 3}
	fmt.Println(arr)

	arr2 := make([]int, 3, 3)
	arr2[0] = 1
	arr2[1] = 2
	arr2[2] = 3
	arr2 = append(arr2, 4, 5, 6)
	fmt.Println(arr2)
	arr2 = append(arr2, arr...)
	fmt.Println(arr2)

	for index, value := range arr2 {
		fmt.Printf("%d-->%d\n", index, value)
	}

	arr3 := []int{1, 2, 3}
	fmt.Println(arr3)
	fmt.Printf("长度是%d，容量是%d", len(arr3), cap(arr3))

	arr3 = append(arr3, 1)
	fmt.Printf("长度是%d，容量是%d", len(arr3), cap(arr3))
}
