package main

import "fmt"

func showMemoryAddress(x *int) {
	fmt.Println(x)
	fmt.Println(*x)
}
func main13() {
	i := 1
	fmt.Println(&i)
	showMemoryAddress(&i)
}
