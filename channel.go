package main

import (
	"fmt"
	"time"
)

//channel 相当于一条管道 一边装水 一边取水
//默认状态下是阻塞的

//功能：1同步  有先有后
// 	   2实现数据交互


//反斜杠双击两下shift 打开快速寻找
//无缓冲channel 管道里面是不能存东西的 作用于两个协程之间的
//在写的时候别人没有读你会阻塞
//管道为空你读数据的时候你会阻塞


//有缓冲channel
//通道满的时候不能放东西
//通道空的时候


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