// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-26 01:48
package main

import "fmt"

func main3() {
	m := new(Movie)
	m.Name = "hello world"
	m.Rating = 3.14
	fmt.Printf("%+v\n", m)
}
