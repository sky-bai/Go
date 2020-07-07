package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word","foo","a string")

	numbPtr := flag.String("numb","42","an int")
	boolPtr := flag.Bool("fork",false,"a bool")

	var svar string
	flag.StringVar(&svar,"svar","bar","a string var")

	flag.Parse()
	//从命令行解析为定义的标志 上面是对于命令行参数的设置

	fmt.Println("word:",*wordPtr)
	fmt.Println("numb:",*numbPtr)
	fmt.Println("bool:",*boolPtr)
	fmt.Println("svar:",svar)
	fmt.Println("tail:",flag.Args())

}
