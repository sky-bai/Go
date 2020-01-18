package main

import (
	"fmt"
	"os"
)

func WriteFile(path string) {
	//新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	//关闭文件  在函数结束之前关闭文件
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		//将字符存入变量 再将变量传入文件
		buf := fmt.Sprintf("bai")
		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
		//返回写入的字节数
	}

}

func Read(path string) {
	//打开文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 1024*2) //2kb
	n, err := f.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)


	//关闭文件
	defer f.Close()
}

func main() {

	//"./" 代表当前路径
	path := "./demo.txt"
	WriteFile(path)

}
