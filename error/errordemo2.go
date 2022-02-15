package error

import (
	"fmt"
	"net"
	"os"
)

func main() {
	/*
		错误的类型
	*/

	//通过error字段来获取
	f, error := os.Open("text.txt")
	if error != nil {
		fmt.Println(error)
		if ins, ok := error.(*os.PathError); ok {
			fmt.Println(ins.Op)
			fmt.Println(ins.Path)
			fmt.Println(ins.Err)
		}
	}
	fmt.Println(f.Name(), "打开文件成功")

	//通过error的方法来获取
	addr, err := net.LookupHost("www.jintiantianqihenhao.com")
	if ins, ok := err.(*net.DNSError); ok {
		if ins.Timeout() {
			fmt.Println("操作超时")
		} else if ins.Temporary() {
			fmt.Println("临时性错误")
		} else {
			fmt.Println("其他错误")
		}
	}
	fmt.Println(err)
	fmt.Println(addr)
}
