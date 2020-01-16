package main

import "fmt"

//底层值为nil
//街边接口内的具体值为nil，方法仍会被nil接受者调用

type Ier interface {
	M()
}

type T2 struct {
	S string
}
//将指针传递就是改变里面的值
func (t *T)M()  {
	if t == nil{
		fmt.Println("nil")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i Ier
	
}