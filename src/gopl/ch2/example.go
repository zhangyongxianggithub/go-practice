package main

import (
	"fmt"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	c := s[len(s)]
	fmt.Println(c)
}
