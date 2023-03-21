package main

import (
	"fmt"
	"strings"
)

func trim(s, cut string, check func(s, cut string) bool, trimFunction func(s, cut string) string) string {
	for check(s, cut) {
		s = trimFunction(s, cut)
	}
	return s
}
func main() {
	fmt.Println(trim("///dsdas///", "/", strings.HasPrefix, strings.TrimPrefix))
}
