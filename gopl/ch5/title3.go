package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil: // 没有宕机
		case bailout{}:
			err := fmt.Errorf("multiple title emlements")
		default:
			panic(p) // 其他类型的宕机，继续宕机
		}
	}()

	// ......
	return "", err
}
