package main

import (
	"fmt"

	"go-practice/gopl/ch7"
)

func main() {
	var w ch7.WordCounter
	w.Write([]byte("hello word go"))
	fmt.Println(w)
}
