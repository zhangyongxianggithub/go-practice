// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-26 01:16
package main

import "fmt"

func main2() {
	cheeses := make([]string, 2)
	cheeses[0] = "hello"
	cheeses[1] = "world"
	cheeses = append(cheeses, "Camembert")
	fmt.Println(cheeses)

}
