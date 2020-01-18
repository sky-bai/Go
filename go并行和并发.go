package main

import (
	"fmt"
	"runtime"
)

//并行 两条队列同时使用两台咖啡机
//并发 两条队列交替使用一台咖啡机
//goroutine 协程 一个协程对应一个任务
//runtime.gosched()让出时间片

func main() {

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println("wangwangwang")
		}
	}()


	for i := 0; i < 5; i++ {
		//让出时间片
		runtime.Gosched()
		fmt.Println("baibaibai")

	}
}
