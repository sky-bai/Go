package main

import "fmt"

func swap(a,b int)  {
	a,b  = b,a
	fmt.Println(a,b)
}

func main()  {
	a,b := 10,20
	//通过一个函数交换a和b的内容
	swap(a,b)  //变量本身传递，值传递
	fmt.Println(a,b)
	//此时main中a，b不会应为swap函数调换
	//这只是将main中a，b的值赋给swap中去

}
