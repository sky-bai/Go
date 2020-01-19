package main

import (
	"fmt"
	"time"
)

//channel 相当于一条管道 一边装水 一边取水
//默认状态下是阻塞的

//功能：1同步  有先有后
// 	   2实现数据交互





func main()  {

	var ch = make(chan string)

	go func() {
		defer fmt.Println("子协程调用完毕")

		for i := 0; i < 2; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}

		ch<-"我是子协程，工作结束"
	}()
	//没有数据就会阻塞
	<-ch
}


//此时什么现象也不会出现 main协程就执行完了 go func协程不会执行