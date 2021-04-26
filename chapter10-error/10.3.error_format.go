//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 01:43
package main

import (
	"fmt"
	"reflect"
)

func main3() {
	name, role := "hello", "world"
	err := fmt.Errorf("the %v %v quit", name, role)
	if err != nil {
		fmt.Println(reflect.TypeOf(err))
		fmt.Println(err)
	}
}
