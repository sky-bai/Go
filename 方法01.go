package main

import "fmt"

//用不同的结构体参数修改成员变量的值
//传递想到参数传递

type Person2 struct {
	name string
	sex  byte
	age  int
}



//setinfovalue 设置结构体成员变量的函数
//用参数传递设置里面的值



//参数为接受者为普通变量 非指针 值语义  就是一份拷贝
func (p Person2) SetinfoValue(n string,s byte,a int)  {
	p.age = a
	//函数是为对象的工具

	p.name = n
	p.sex = s
	fmt.Println(&p)
}

//参数为指针变量 引用传递
func (p *Person2) SetinfoPoint(n string,s byte,a int)  {
	p.age = a
	p.name = n
	p.sex = s
	fmt.Println(p)

	//将要改变的结构体的地址传过来
}


func main() {
	var s1 Person2
	fmt.Println(s1)
	fmt.Println(&s1)

	s1.SetinfoValue("lin",'o',20)
	fmt.Println(s1)

	s1.SetinfoPoint("bai",'s',18)
	fmt.Println(s1)



	//修改一个结构体变量的值要用引用传递
	//如果单单使用函数进行修改 只是对原有结构体进行一份拷贝


	//如何改变结构体的值
	//将结构体变为函数的接受者
	//将结构体的指针传过来 我再改变它的值

	//方法是为改变而生的









}