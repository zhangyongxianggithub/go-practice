package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(0)
	n = runtime.GOMAXPROCS(0)
	fmt.Println(n)
}
