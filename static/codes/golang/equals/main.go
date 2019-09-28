package main

import "fmt"

type Foo struct {
	Value string
}

func (f Foo) Interface() (i interface{}) {
	return
}

func main() {
	foo1 := Foo{
		Value: "sdf",
	}
	foo2 := Foo{
		Value: "abc",
	}
	fmt.Println(foo1 == foo2)
}
