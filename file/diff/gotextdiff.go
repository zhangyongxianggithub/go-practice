package main

import (
	"fmt"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

func main() {
	edits := myers.ComputeEdits(span.URIFromPath("a.txt"), "aaaaaa", "bbb")
	unified := gotextdiff.ToUnified("a.txt", "b.txt", "aaaaaa", edits)

	diff := fmt.Sprint(unified)
	fmt.Println(diff)
}
