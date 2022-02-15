package _type

import "time"

func main() {
}

type MyDuration time.Duration

//type MyDuration = time.Duration//这样会报错, 因为如果使用=, 其实和Duration是一回事儿, Duration不能新增方法
func (d MyDuration) SipleSet() {

}
