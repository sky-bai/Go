package main
//接口是一组方法签名定义的集合
//描述一些方法的集合
//接口习惯以er结尾
//接口只有方法声明没有实现 没有数据字段
//接口可以匿名嵌入其他接口，或者嵌入其他接口中
import "fmt"

type humaner interface {
	//方法只有声明没有实现
	sayhi()
}

type Student2 struct {
	name string
	id int
}

func (tem *Student2)sayhi()  {
fmt.Println(tem.name,tem.id)
}


type Teacher struct {
	addr string
	group int
}

func (tem *Teacher)sayhi()  {
	fmt.Println(tem.addr,tem.group)
}


//封装  通过方法
//继承 通过匿名字段
//多态 通过接口实现

func main()  {
	//定义接口类型的变量
	var i humaner

	//只要实现此接口方法类型，那么这个类型的变量(接受者类型)就可以给i赋值
	s := &Student2{"bai",9}
	i := s
	i.sayhi() //根据前面所实例化的对象 自动匹配它所选择的方法

	//将方法接受者的类型变量赋给接口变量 使接口变量实现该方法
	w := &Teacher{"bai",8}
	i = w
	i.sayhi()









}
//结构体与方法相结合