//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-26 01:40
package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

func main1() {
	m := Movie{
		Name:   "Citizen Kane",
		Rating: 3.14,
	}
	fmt.Println(m)
}
