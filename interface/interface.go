package _interface

import "fmt"

func main() {
	/*
			接口: interface
				在go中, 接口是一组方法的签名

				当某个类型为这个接口中的所有方法提供了方法的实现, 它别称为接口

		1,当需要使用接口类型的对象时, 可以使用任意实现累代替
		2,接口对象不能访问实现类的字段

		在go语言中接口是非侵入型的
		go语言中通过接口来模拟多态

		接口的用法:
			1, 一个函数如果接受接口类型的参数, 那么实际上就可以传入该接口的任意实现类型对象作为参数
			2, 定义一个类型为接口类型, 实际上可以赋值为任意实现类的对象
	*/
	//1, 创建Mouse类型
	m1 := Mouse{"鼠标"}
	fmt.Println(m1.name)
	//2, 创建FlashDisk类型
	f2 := FlashDisk{"键盘"}
	fmt.Println(f2.name)

	testInterface(m1)
	testInterface(f2)

	var usb USB
	usb = m1
	usb.start()
	usb.end()
}

//1,定义接口
type USB interface {
	start()
	end()
}

//2.定义实现类
type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println("鼠标插入")
}

func (m Mouse) end() {
	fmt.Println("鼠标拔出")
}

func (f FlashDisk) start() {
	fmt.Println("键盘插入")
}

func (f FlashDisk) end() {
	fmt.Println("键盘拔出")
}

//测试方法
func testInterface(u USB) {
	u.start()
	u.end()
}
