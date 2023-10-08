// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-26 02:13
package main

import "fmt"

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := &a
	b.Ice = false
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b)
}
