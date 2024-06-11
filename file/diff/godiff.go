package main

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	text1 = `Lorem ipsum dolor
ssssss
	qeweq
dsadas
dsadsad`
	text2 = `Lorem dolor sit amet.
dsada
dsadsfv
54t5`
)

func main() {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(text1, text2, true)

	fmt.Println(dmp.DiffPrettyText(diffs))
	for _, d := range diffs {
		fmt.Println(d.Text)
	}
}
