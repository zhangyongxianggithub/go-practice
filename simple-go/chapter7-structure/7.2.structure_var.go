// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-26 01:44
package main

import "fmt"

func main2() {
	var m Movie
	fmt.Printf("%+v\n", m)
	m.Name = "Metropolis"
	m.Rating = 3.14
	fmt.Printf("%+v\n", m)
}
