package main

import "fmt"

//匿名函数定义

func main() {
	A := 10
	str := 20


	//匿名函数 通过一个变量来接收 没有函数名字
	//匿名函数形成一个闭包 可以访问闭包以外的作用域
	f1 := func() {
		fmt.Println(A)
		fmt.Println(str)
	}

	//调用匿名函数   变量加圆括号
	f1()

	//匿名函数加调用
	func (){
		fmt.Println(A)
	}()   //后面的括号为调用  里面是参数




	//






}
