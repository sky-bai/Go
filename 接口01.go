package main

import "math"

//接口是一组方法签名定义的集合
//描述一些方法的集合
//接口习惯以er结尾
//接口只有方法声明没有实现 没有数据字段
//接口可以匿名嵌入其他接口，或者嵌入其他接口中


type MyFloat1 float64

func (f MyFloat1)Abs() float64 {
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}
type Vertex1 struct {
	X, Y float64
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//-------------------------------------------------------------------------------

type Abser interface {
	Abs() float64
}

func main()  {
	var a Abser
	f :=MyFloat1(-math.Sqrt2)
	v :=Vertex1{3,4}

	a = f
	a = &v


}
