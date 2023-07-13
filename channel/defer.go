package main

import (
	"fmt"
	"net/url"
)

func main() {
	defer func() {
		fmt.Println("first defer")
	}()
	defer func() {
		fmt.Println("second defer")
	}()
	defer func() {
		fmt.Println("third defer")
	}()
	fmt.Println("main")

	fmt.Println(url.QueryEscape("/dadasdas/dsadas张"))
}
