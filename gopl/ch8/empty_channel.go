package main

import (
	"fmt"
	"time"
)

func main() {
	// test := make(chan int, 1)
	test := make(chan int)
	// var test1 chan int
	go func() {
		time.Sleep(2 * time.Second)
		test <- 1
	}()
	fmt.Println(<-test)

}
