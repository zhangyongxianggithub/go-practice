package main

import (
	"fmt"

	"gopl/ch7/mod"
)

func main() {
	var w mod.WordCounter
	w.Write([]byte("hello word go"))
	fmt.Println(w)
}
