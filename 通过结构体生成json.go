package main

import (
	"encoding/json"
	"fmt"
)

//json在服务器和客户端之间的一种数据交换格式
//json是键值对集合
//要生成json 成员变量名必须大写

type IT struct {
	Company  string
	Subjects []string //使用切片类型
	Isok     bool
	Price    float64
}

func main() {
	s := IT{
		Company:  "BAT",
		Subjects: []string{"go", "python"},
		Isok:     true,
		Price:    210,
	}

	//根据内容生成json文本 文本
	buf, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
