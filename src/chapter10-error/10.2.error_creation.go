//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 01:41
package main

import (
	"errors"
	"fmt"
)

func main2() {
	err := errors.New("something went wrong")
	if err != nil {
		fmt.Println(err)
	}
}
