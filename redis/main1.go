package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go向redis写入数据和读取数据
	//1,连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
	}
	defer conn.Close()
	//2,写入数据
	_, err = conn.Do("HMSet", "user", "name", "Tom", "age", 19)
	if err != nil {
		fmt.Println("redis set error:", err)
		return
	}
	//3,获取数据
	r, err := redis.Strings(conn.Do("HMGet", "user", "name", "age"))
	if err != nil {
		fmt.Println("redis get error:", err)
		return
	}
	fmt.Printf("r:%v\n", r)
	//返回的r是interface，需要根据对应的值转换
	//for index, value := range r {
	//	fmt.Printf("%d:%v\n", index, value)
	//}
}
