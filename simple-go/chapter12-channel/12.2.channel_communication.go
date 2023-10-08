// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-27 02:14
package main

import (
	"fmt"
	"time"
)

func slowFunc2(c chan string) {
	time.Sleep(time.Second * 2)
	fmt.Println("sleep finished")
	c <- "slow function complete"
}
func main2() {
	c := make(chan string)
	go slowFunc2(c)
	msg := <-c
	fmt.Println(msg)

}
