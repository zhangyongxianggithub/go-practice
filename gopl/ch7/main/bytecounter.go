package main

import (
	"fmt"

	"go-practice/gopl/ch7"
)

func main() {
    var c ch7.ByteCounter
    c.Write([]byte("hello"))
    fmt.Println(c)
    c = 0
    var name = "Dolly"
    fmt.Fprintf(&c, "hello,%s", name)
    fmt.Println(c)
}
