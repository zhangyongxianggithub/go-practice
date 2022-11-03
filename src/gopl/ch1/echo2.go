package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		s += sep + string(index) + "," + arg
		sep = " "
	}
	fmt.Println(s)
}
