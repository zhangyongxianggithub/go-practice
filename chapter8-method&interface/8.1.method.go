//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-26 02:26
package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}

func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	fmt.Println(r)
	return m.Name + ", " + r
}
func main1() {
	m := Movie{
		Name:   "Spiderman",
		Rating: 3.2,
	}
	fmt.Println(m.summary())
}
