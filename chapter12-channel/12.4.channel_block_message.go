//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 02:32
package main

import (
	"fmt"
	"time"
)

func slowFunc4(c chan string) {
	t := time.NewTicker(time.Second * 1)
	for {
		c <- "ping"
		<-t.C
	}
}
func main4() {
	messages := make(chan string)
	go slowFunc4(messages)
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}
