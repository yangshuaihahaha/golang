package oita

func main() {
	//1，iota只能在常量的表达式中使用。
	//2，每次 const 出现时，都会让 iota 初始化为0.
	const a = iota // a=0
	const (
		b = iota //b=0
		c        //c=1
	)
	//3，自增长常量经常包含一个自定义枚举类型，允许你依靠编译器完成自增设置。
	type Stereotype int

	const (
		TypicalNoob           Stereotype = iota // 0
		TypicalHipster                          // 1
		TypicalUnixWizard                       // 2
		TypicalStartupFounder                   // 3
	)
	//4，可跳过的值
	//我们可以使用下划线跳过不想要的值。
	type AudioOutput int

	const (
		OutMute   AudioOutput = iota // 0
		OutMono                      // 1
		OutStereo                    // 2
		_
		_
		OutSurround // 5
	)
	//5，位掩码表达式
	type Allergen int

	const (
		IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
		IgChocolate                         // 1 << 1 which is 00000010
		IgNuts                              // 1 << 2 which is 00000100
		IgStrawberries                      // 1 << 3 which is 00001000
		IgShellfish                         // 1 << 4 which is 00010000
	)
	//6，定义数量级
	type ByteSize float64

	const (
		_           = iota             // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
		MB                             // 1 << (10*2)
		GB                             // 1 << (10*3)
		TB                             // 1 << (10*4)
		PB                             // 1 << (10*5)
		EB                             // 1 << (10*6)
		ZB                             // 1 << (10*7)
		YB                             // 1 << (10*8)
	)
	//7，定义在一行的情况
	const (
		Apple, Banana = iota + 1, iota + 2
		Cherimoya, Durian
		Elderberry, Fig
	)
	//iota 在下一行增长，而不是立即取得它的引用。
	// Apple: 1
	// Banana: 2
	// Cherimoya: 2
	// Durian: 3
	// Elderberry: 3
	// Fig: 4
	//8，中间插队
	const (
		i = iota
		j = 3.14
		k = iota
		l
	)
	//那么打印出来的结果是 i=0,j=3.14,k=2,l=3
}
