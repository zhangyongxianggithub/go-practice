//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 00:38
package main

import "fmt"

type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) area() float64 {
	return 0.5 * t.height * t.base
}
func main3() {
	t := Triangle{
		base:   3,
		height: 1,
	}
	fmt.Println(t.area())
}
