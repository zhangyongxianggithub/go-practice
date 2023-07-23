package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "aaaaxlsx."
	index := strings.LastIndex(a, ".")
	if index >= 0 {
		fmt.Println(index)
		fmt.Println(a[:index])
		fmt.Println(a[index:])
	}

}
