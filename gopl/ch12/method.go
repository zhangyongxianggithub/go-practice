package main

import (
	"fmt"
	"reflect"
)

func test() {
	fmt.Println("test")
}

func test1() {
	fmt.Println("test1")
}

type Func func()

func main() {
	var f Func = test1
	fmt.Println(reflect.TypeOf(f).Name())
}
