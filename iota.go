package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c
		d = "haha"
		e
		f = 100
		g
		h = iota
		i
	)
	//var a1 int32 = 31
	//var a2 float32 = 33

	//a1 = int32(a2)
	fmt.Println(a, b)

}
