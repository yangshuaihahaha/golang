package main

import (
	"bufio"
	"fmt"
	"os"
)

type ChartCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	Other      int
}

func main() {
	//统计英文，数字，空格和其他字符的数量
	srcFileName := "/Users/yangshuai/go/src/demo3/filedemo/abc.text"

	file, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s\n", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
