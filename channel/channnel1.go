package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := make(chan string, 10)
	go func() {
		for {
			if m, ok := <-c; ok {
				fmt.Println("goroutine 1 message: " + m)
			} else {
				fmt.Printf("goroutine 1 end: %s\n", m)
				return
			}
		}
	}()
	go func() {
		for {
			if m, ok := <-c; ok {
				fmt.Println("goroutine 2 message: " + m)
			} else {
				fmt.Printf("goroutine 2 end: %s\n", m)
				return
			}
		}
	}()
	for i := 0; i < 10; i++ {
		c <- "message:" + strconv.Itoa(i)
	}
	close(c)

	time.Sleep(1 * time.Second)

}
