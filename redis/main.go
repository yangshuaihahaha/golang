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
	_, err = conn.Do("Set", "name", "Tom")
	if err != nil {
		fmt.Println("redis set error:", err)
		return
	}
	//3,获取数据
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("redis get error:", err)
		return
	}
	//返回的r是interface，需要根据对应的值转换
	fmt.Println(r)

}
