package main

import (
	"fmt"
)

type func1 func()

func (f func1) Run() { f() }

func test(funcParameter func()) {
	f := func1(funcParameter)
	f.Run()
}

func main() {
	test(func() {
		fmt.Println("为什么这么写呢")
	})
}
