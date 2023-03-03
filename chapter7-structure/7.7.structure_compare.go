//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-26 02:03
package main

import "fmt"

type Drink struct {
	Name string
	Ice  bool
}

func main7() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	if a == b {
		fmt.Println("a and b are the same")
	}
}
