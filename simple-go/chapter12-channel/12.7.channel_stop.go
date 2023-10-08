// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-27 02:48
package main

import (
	"fmt"
	"time"
)

func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I am sending a message"
		<-t.C
	}
}
func main() {
	messages := make(chan string)
	stop := make(chan bool)
	go sender(messages)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("now stop")
		stop <- true
	}()
	for {
		select {
		case <-stop:
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}
