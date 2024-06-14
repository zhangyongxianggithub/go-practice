package main

import (
	"fmt"
)

var m map[string]Test

func main() {
	var m = make(map[string]Test)
	t1 := &Test1{Name: "test1"}
	t2 := &Test2{Name: "test2"}
	m["1"] = t1
	m["2"] = t2
	for k, v := range m {
		fmt.Println("k", k, ",v", v)
	}
	t1.Name = "test1-update"
	for k, v := range m {
		fmt.Println("k", k, ",v", v)
	}
}

type Test interface {
	fmt.Stringer
}

type Test1 struct {
	Name string
}

func (t *Test1) String() string {
	return t.Name
}

type Test2 struct {
	Name string
}

func (t *Test2) String() string {
	return t.Name
}
