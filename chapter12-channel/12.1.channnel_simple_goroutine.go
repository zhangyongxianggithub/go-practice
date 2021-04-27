//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 02:11
package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleep finished")
}
func main1()  {
	go slowFunc()
	fmt.Println("I am not show until slowFunc() completes")
	time.Sleep(time.Second * 4)
}
