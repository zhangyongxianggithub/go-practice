package main

import "fmt"

var s10 = "hello world"

func main10() {
	fmt.Printf("Print 's' variable from outer block %v\n", s10)
	b := true
	if b {
		fmt.Printf("Print 'b' variable from outer block %v\n", b)
		i := 42
		if b != false {
			fmt.Printf("Print 'i' variable from outer block %v\n", i)
		}
	}
}
