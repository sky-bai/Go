package main

import "fmt"

//指针做函数参数
//交换变量地址

func swap(p1, p2 *int) {
	p1,p2 = p2,p1

}
func main() {
	a,b := 10,20
	a,b = b,a
	swap(&a,&b)
	fmt.Println(a,b)
}
