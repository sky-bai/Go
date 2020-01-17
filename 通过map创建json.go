package main

import (
	"encoding/json"
	"fmt"
)

//创建一个map
func main() {
	//v interfaca{} 空接口 万能参数
	m := make(map[string]interface{}, 4)
	m["company"] = "itcase"
	//注意这里里面用切片
	m["subjects"] = []string{"go", "python"}
	m["isok"] = true
	m["price"] = 666

	r, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	//打印的时候一定要注意要json。marshal返回的是字节切片 要强转为string
	fmt.Println(string(r))
}
