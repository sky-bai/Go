package main

import (
	"fmt"
	"math"
)

//为非结构体类型声明方法
type MyFloat float64

func (f MyFloat)Abs()float64  {
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}


func main()  {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
//方法是为某一类型而生的权力  为结构体 为数据类型 对某一数据进行判断