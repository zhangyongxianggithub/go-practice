package main

import (
	"fmt"
)

func main() {
	var a [3]int             // 声明3个整数的数组
	fmt.Println(a[0])        // 输出数组的第一个元素
	fmt.Println(a[len(a)-1]) // 输出数组的最后一个元素
	for i, v := range a {
		fmt.Printf("%d, %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q = [3]int{1, 2, 3}
	fmt.Println(q[2])
	t := [...]int{1, 2, 3} // 通过初始化的值决定数组的长度
	fmt.Println(t)
	fmt.Println(symbol)
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

var symbol = [...]string{USD: "$", EUR: "e", GBP: "r", RMB: "¥"}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
