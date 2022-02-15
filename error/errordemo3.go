package error

import (
	"fmt"
	"math"
)

func main() {
	/*
		自定义错误
	*/

	area, err := circleArea(-3.2)
	fmt.Println(err)
	if err != nil {
		if ins, ok := err.(*areaError); ok {
			fmt.Println(ins.msg)
		}
	}
	fmt.Println(area)
}

//1, 定义一个结构体, 表示错误的类型
type areaError struct {
	msg    string
	radius float64
}

//2,实现error的接口, 其实就是实现error的方法
func (e *areaError) Error() string {
	return fmt.Sprintf("error: 半径，%.2f，%s", e.radius, e.msg)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"半径是非法的", radius}
	}
	return math.Pi * radius * radius, nil
}
