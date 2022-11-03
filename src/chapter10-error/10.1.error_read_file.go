//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 01:37
package main

import (
	"fmt"
	"io/ioutil"
)

func main1() {
	file, err := ioutil.ReadFile("foo.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", file)
	}
}
