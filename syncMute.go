package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

// singleHandler 代表单次处理函数的类型。
type singleHandler func() (data string, n int, err error)

type handlerConfig struct {
	handler   singleHandler //单次处理函数
	goNum     int           // 需要启用的goroutine的数量。
	number    int           // 单个goroutine中的处理次数。
	interval  time.Duration //单个goroutine处理的间隔时间
	counter   int
	counterMU sync.Mutex //数据量计数器专用的互斥锁
}

//会增加计算器的值 并返回计数器的个数
func (hc *handlerConfig) count(increment int) int {
	hc.counterMU.Lock()
	defer hc.counterMU.Unlock()
	hc.counter += increment
	return hc.counter
}

func main() {
	var mu sync.Mutex
	//用于生成写入函数的函数
	genWriter := func(writer io.Writer) singleHandler {
		return func() (data string, n int, err error) {
			//准备数据
			data := fmt.Sprintf("%s\t",
				time.Now().Format(time.StampNano))
			//写入数据
			mu.Lock()
			defer mu.Unlock()
			n, err = writer.Write([]byte(data))
			return
		}

	}
	//用于生成读入函数的函数
	genReader := func(reader io.Reader) singleHandler {
		return func() (data string, n int, err error) {
			buffer, ok := reader.(*bytes.Buffer)
			if !ok {
				err = errors.New("unsupported reader")
				return
			}
			// 读取数据。
			mu.Lock()
			defer mu.Unlock()
			data, err = buffer.ReadString('\t')
			n = len(data)
			return
		}
	}

	var buffer bytes.Buffer

	//数据写入配置
	writingConfig := handlerConfig{
		handler:  genWriter(&buffer),
		goNum:    5,
		number:   4,
		interval: time.Millisecond * 100,
	}

	// 数据读取配置。
	readingConfig := handlerConfig{
		handler:  genReader(&buffer),
		goNum:    10,
		number:   2,
		interval: time.Millisecond * 100,
	}

	//代表信号的通道
	sign := make(chan struct{}, writingConfig.goNum+readingConfig.goNum)

	// 启用多个goroutine对缓冲区进行多次数据写入。
	for i := 1; i <= writingConfig.goNum; i++ {
		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()
			// 一个goroutine中的处理次数。
			for j := 1; i <= writingConfig.number; j++ {
				time.Sleep(writingConfig.interval) //单个goroutine处理的间隔时间
				data, n, err := writingConfig.handler()
				if err != nil {
					log.Printf("writer [%d-%d]: error: %s",
						i, j, err)
					continue
				}
			}
		}(i)
	}

}
//为了保证每次读写的同步 每个goroutine对单一变量的操作