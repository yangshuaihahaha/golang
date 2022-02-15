package error

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, error := os.Open("text.txt")
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(f.Name(), "打开文件成功")

	/*
		error: 内置的数据类型, 内置的接口
			定义方法: Error() string
		使用go语言提供好的包
			errors包下的函数: New(), 创建一个error对象
			func包下的Errorf()函数
				func Errorf(format string, a ...interface{}) error
	*/
	//1,创建一个error数据
	err1 := errors.New("自己创建的error数据")
	fmt.Println(err1)
	fmt.Printf("%T\n", err1)

	//2,另一个创建error的方法
	err2 := fmt.Errorf("自己创建的另一个error数据")
	fmt.Println(err2)
	fmt.Printf("%T\n", err2)

	err3 := checkAge(-1)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println("程序继续。。。")
}

//设计一个函数, 如果不合法就返回error
func checkAge(age int) error {
	if age < 0 {
		//return errors.New("年龄不合法")
		return fmt.Errorf("年龄不合法")
	}
	fmt.Println("年龄是:", age)
	return nil
}
