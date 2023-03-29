package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("children goroutine!!!")
	}()
	// time.Sleep(1 * time.Second)
	fmt.Println("main goroutine")
}
