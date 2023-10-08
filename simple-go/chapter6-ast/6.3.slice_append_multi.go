// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-26 01:18
package main

import "fmt"

func main3() {
	cheeses := make([]string, 2)
	cheeses[0] = "hello"
	cheeses[1] = "world"
	appendedSlice := append(cheeses, "Camembert", "Reblochon", "Picodon")
	fmt.Println(cheeses)
	fmt.Println(appendedSlice)
}
