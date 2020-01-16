package main

import "fmt"

type i interface {
	M()
}

type T struct {
	S string
}

func (t T)M()  {
	fmt.Println(t.S)
	
}

func main() {
	var q i = T{"s"}
	q.M()
}