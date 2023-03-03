//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 01:59
package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleep finished")
}
func main1() {
	slowFunc()
	fmt.Println("I am not show until slowFunc() completes")
}
