package main

import (
	"fmt"
	"reflect"
)

func main4() {
	var s = "string"
	var i = 10
	var f float32 = 0.12
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
}
