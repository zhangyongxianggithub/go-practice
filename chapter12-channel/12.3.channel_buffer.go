//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 02:23
package main

import (
	"fmt"
	"time"
)

func slowFunc3(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}
func main3() {
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"
	close(messages)
	fmt.Println("push two messages into channel")
	time.Sleep(time.Second * 1)
	slowFunc3(messages)

}
