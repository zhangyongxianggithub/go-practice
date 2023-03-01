package main

import (
	"fmt"
)

func main() {
	var runes []rune
	for _, r := range "Hello, world" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
	var runes2 = []rune("Hello, world")
	fmt.Printf("%q\n", runes2)

}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z

}
