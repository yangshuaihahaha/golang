package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大的空闲连接数
		MaxActive:   0,   //表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 100, //最大的空闲时间
		Dial: func() (conn redis.Conn, e error) { //初始化连接的代码
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	//从pool取出一个连接
	conn := pool.Get()
	defer conn.Close()
	//2,写入数据
	_, err := conn.Do("Set", "name", "Tom猫")
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
	fmt.Println("r:", r)
	//连接池关闭后就不能再取出东西了
}
