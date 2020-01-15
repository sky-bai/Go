package main

import "fmt"

//Student1中包含Person1同名字段

type Person1 struct {
	name string
	sex byte
	age int
}


type Student1 struct {
	Person  //匿名字段   只有类型没有名字  继承了person的成员
	id int
	addr string
	name string
}

func main()  {

	var s Student1
	s.Person1.name = "mike"
	fmt.Println(s)
}