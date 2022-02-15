package main

import (
	"fmt"
)

/*
	main函数和init函数
	相同点：
	两个函数定义时不能有任何参数和返回值
	该函数只能由go程序自动调用，不可以被引用
	不同点：
	init可以应用于任意的包中，且可以重复定义多个。
	main函数只能用于main包中，且只能定义一个
	两个函数的执行顺序
	在main包中的go文件总是会被执行
	对同一个go文件的init()调用顺序是从上到下的。
	对于同一个package中的不同文件，将文件名字按字符串进行"从小到大"排序，之后调用各文件中的init函数
	对于不同的package，如果不相互依赖的话，按照main包中的import的顺序调用包中的init()函数
	如果package存在依赖，调用顺序为最后被依赖的最先执行，例如：导入顺序main -> A --> B --> C
	则初始化顺序为C --> B --> A --> main，一次执行对应的init方法。main包总是被最后一个初始化，因为他总是依赖别的包
*/

func main() {
	MyTest()
	fmt.Println("这是main函数")
}

func init() {
	fmt.Println("main中的init函数")
}
