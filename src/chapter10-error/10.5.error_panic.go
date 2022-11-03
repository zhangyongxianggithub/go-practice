//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 01:52
package main

import "fmt"

func main() {
	fmt.Println("hello")
	panic("Oh no, Goodbye")
	fmt.Println("not executed")
}
