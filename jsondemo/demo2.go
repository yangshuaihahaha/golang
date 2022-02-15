package main

import (
	"encoding/json"
	"fmt"
)

type Monster1 struct {
	Name     string `json:"monster_name"` //更改序列化时候的别名
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func main() {
	body := "{\"monster_name\":\"小丑\",\"Age\":12,\"Birthday\":\"2007-01-01\",\"Sal\":12,\"Skill\":\"喷火\"}"
	var monster Monster1
	err := json.Unmarshal([]byte(body), &monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(monster)
}
