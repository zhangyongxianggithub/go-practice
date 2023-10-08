// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-26 01:53
package main

import "fmt"

type Address struct {
	Number int
	Street string
	City   string
}
type Superhero struct {
	Name    string
	Age     int
	Address Address
}

func main5() {
	e := Superhero{
		Name: "Batman",
		Age:  123,
		Address: Address{
			Number: 1007,
			Street: "Mountain Drive",
			City:   "Gotham",
		},
	}
	fmt.Printf("%+v\n", e)
}
