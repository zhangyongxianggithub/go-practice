// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-26 01:02
package main

import "fmt"

func main() {
	defer fmt.Println("I am the first defer statement")
	defer fmt.Println("I am the second defer statement")
	defer fmt.Println("I am the third defer statement")
	fmt.Println("Hello defer")
}
