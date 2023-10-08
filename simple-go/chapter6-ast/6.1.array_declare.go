// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-26 01:11
package main

import "fmt"

func main1() {
	var cheeses [2]string
	fmt.Println(cheeses)
	cheeses[0] = "hello"
	cheeses[1] = "world"
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	fmt.Println(cheeses)
}
