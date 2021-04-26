//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 00:41
package main

import "fmt"

func (t Triangle) changeBase(f float64) {
	t.base = f
}
func main4() {
	t := Triangle{
		base:   3,
		height: 1,
	}
	t.changeBase(4)
	fmt.Println(t.base)
}
