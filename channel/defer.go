package main

import (
	"fmt"
	"io"
	"net/url"
)

func main() {
	var r io.Reader
	defer func() {
		r.Read(make([]byte, 1024))
		fmt.Println("first defer")
	}()
	fmt.Println("main")
	defer func() {
		fmt.Println("second defer")
	}()
	defer func() {
		fmt.Println("third defer")
	}()
	fmt.Println("main")

	fmt.Println(url.QueryEscape("/dadasdas/dsadas张"))
}
