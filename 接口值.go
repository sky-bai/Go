package main

import "fmt"

//接口值
//接口也是值 ，它可以像值一样传递
//接口是看成是包含值和类型的元组  (value,type)

type I1 interface {
	M()
}


type T1 struct {
	S string
}

func (t *T1)M()  {//该方法用指针作为接受者
	fmt.Println(t.S)//打印该结构体里面的值
}

type F float64

func (f F) M()  {
	fmt.Println(f)
}


func main()  {
	//x先定义I接口变量
	var i1 I1
	//赋值 如果是结构体就取地址符 还要初始化结构体中变量
	i1 = &T1{"bai"}
	i1.M()

	i1 = F(56)
	i1.M()

}
