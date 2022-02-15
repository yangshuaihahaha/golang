package main

import (
	"fmt"
	"time"
)

func main() {
	//获取现在的时间
	t1 := time.Now()
	fmt.Println(t1)

	//获取当前的时间
	t2 := time.Date(2009, 7, 12, 13, 34, 56, 0, time.Local)
	fmt.Println(t2)

	//time转字符串
	//2006年1月2日 15:04:05 模版数据是固定的,格式可以随意写
	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s1)

	//字符串转time
	s2 := "2019-9-10"
	t3, error := time.Parse("2006-1-2", s2)
	if error == nil {
		fmt.Println("t3:", t3)
	} else {
		fmt.Println(error)
	}

	//根据当前时间获取指定的内容
	year, month, day := t1.Date()
	fmt.Println(year, month, day)

	hour, min, sec := t1.Clock()
	fmt.Println(hour, min, sec)

	fmt.Println(t1.Year())
	fmt.Println(t1.Month())
	fmt.Println(t1.Day())
	fmt.Println(t1.Hour())
	fmt.Println(t1.Minute())
	fmt.Println(t1.Second())

	//获取时间戳
	t4 := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	timeStamp1 := t4.Unix() //秒的差值
	fmt.Println("timeStamp1:", timeStamp1)

	fmt.Println(t1.Unix())     //秒时间戳
	fmt.Println(t1.UnixNano()) //纳秒时间戳

	//时间的间隔
	t5 := t1.Add(time.Minute)
	fmt.Println(t1)
	fmt.Println(t5)

	fmt.Println(t1.Add(time.Hour * 24))
	fmt.Println(t1.AddDate(1, 0, 0))

	d1 := t5.Sub(t1)
	fmt.Println(d1)

	//睡眠
	var s3 int
	s3 = 6
	time.Sleep(time.Second * time.Duration(s3))
	fmt.Println("主函数结束")

}
