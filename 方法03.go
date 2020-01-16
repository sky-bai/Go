package main

import (
	"fmt"
	"math"
)

//为指针接受者声明方法
type Vertex struct {
	x,y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x)
}
func (v *Vertex) Scale(f float64)  {
	v.x = v.x * f
	v.y = v.y * f
}

func main()  {
	v := Vertex{x:3.14,y:3.14}
	fmt.Println(v.Abs())
	v.Scale(3)
}
//每个结构体都应该有值接受者和指针接受者



