package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"monster_name"` //更改序列化时候的别名
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func main() {
	//序列化结构体
	monster := Monster{"小丑", 12, "2007-01-01", 12, "喷火"}
	data, error := json.Marshal(&monster)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(data))

	//序列化map
	a := make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "hongyadong"
	data, error := json.Marshal(a)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(data))

	//序列化slice
	var slice []map[string]interface{}
	a := make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "hongyadong"
	slice = append(slice, a)
	b := make(map[string]interface{})
	b["name"] = "牛魔王"
	b["age"] = 60
	b["address"] = "火焰山"
	slice = append(slice, b)
	data, error := json.Marshal(slice)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(data))

	//对基本数据类型序列化(没有什么意义)
	num := 23.78
	data, error := json.Marshal(num)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(data))

}
