package main

import (
	"fmt"

	"go-practice/gopl/ch7/mod"
)

func main() {
	// 计算单词数
	var w mod.WordCounter
	w.Write([]byte("hello word go"))
	fmt.Println(w)
	// 计算行数
	var l mod.LineCounter
	l.Write([]byte("hello\nword\ngo\n\n"))
	fmt.Println(l)
}
