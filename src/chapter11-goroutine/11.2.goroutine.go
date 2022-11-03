//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 02:02
package main

import (
	"fmt"
	"time"
)

func main() {
	go slowFunc()
	fmt.Println("I am not show until slowFunc() completes")
	time.Sleep(time.Second * 4)
}
