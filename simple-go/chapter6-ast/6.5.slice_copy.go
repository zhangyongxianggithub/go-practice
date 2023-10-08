// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-26 01:27
package main

import "fmt"

func main5() {
	cheeses := make([]string, 3)
	cheeses[0] = "hello"
	cheeses[1] = "world"
	smellyCheeses := make([]string, 2)
	copy(smellyCheeses, cheeses[1:])
	fmt.Println(len(smellyCheeses))
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
	fmt.Println(smellyCheeses)
}
