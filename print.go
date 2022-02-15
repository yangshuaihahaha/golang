package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("ha--%v--ha\n", 12)
	fmt.Printf("ha--%T--ha\n", 12)
	fmt.Printf("ha--%t--ha\n", true)
	fmt.Printf("ha--%s--ha\n", "123")
	fmt.Printf("ha--%f--ha\n", 3.14)
	fmt.Printf("ha--%d--ha\n", 31)
	fmt.Printf("ha--%b--ha\n", 31)
	fmt.Printf("ha--%o--ha\n", 31)
	fmt.Printf("ha--%X--ha\n", 31)
	fmt.Printf("ha--%x--ha\n", 31)
	f := 'A'
	fmt.Printf("ha--%c--ha\n", f)
	fmt.Printf("ha--%p--ha\n", &f)

	fmt.Printf("--------------------------\n")

	//var x int
	//var y float64
	//fmt.Println("请输入x，y")
	//fmt.Scanln(&x, &y)
	//fmt.Printf("x的数值是：%d, y的数值是%f",x,y)

	fmt.Println("请输入一个字符：")
	reader := bufio.NewReader(os.Stdin)
	sq, _ := reader.ReadString('\n')
	fmt.Println("读到的数据是：", sq)
}
