// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-26 01:22
package main

import "fmt"

func main4() {
	cheeses := make([]string, 3)
	cheeses[0] = "Hello"
	cheeses[1] = "world"
	cheeses[2] = "Camembert"
	fmt.Println(cheeses)
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println(cheeses)

}
