package main

import (
	"fmt"
	"time"
)

func main() {

	a := 1

	go func() {
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("recover %s", p)
			} else {
				fmt.Println("no exception")
			}

		}()
		if a == 0 {
			panic("ddddddddd")
		}

	}()
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("recover xxxxx%s \n", p)
		} else {
			fmt.Println("main no exception")
		}
	}()
	fmt.Println(`execute`)
	panic("main goroutine")
	time.Sleep(1 * time.Minute)

}
