package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	fmt.Println("initfile")
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

/*种群统计，这里是怎么算的呢*/
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] + pc[byte(x>>(1*8))] + pc[byte(x>>(2*8))] + pc[byte(x>>(3*8))] + pc[byte(x>>(4*8))] + pc[byte(x>>(5*8))] + pc[byte(x>>(6*8))] + pc[byte(x>>(7*8))])
}

func main() {
	for i, v := range pc {
		fmt.Println(i, v)
	}
}
