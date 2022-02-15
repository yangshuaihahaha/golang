package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	num1 := rand.Int()
	fmt.Println(num1)

	num1 = rand.Intn(10)
	fmt.Println(num1)
	//设置种子数
	rand.Seed(time.Now().Unix())
	num1 = rand.Intn(10)
	fmt.Println(num1)

	num3 := rand.Intn(8) + 3
	fmt.Println(num3)
}
