package main

import (
	"fmt"

	"gopl/ch7/mod"
)

func main() {
	var c mod.ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello,%s", name)
	fmt.Println(c)
}
