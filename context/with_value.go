package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	child := context.WithValue(ctx, "aaa", 1)
	fmt.Println(ctx.Value("aaa"))

	fmt.Println(child.Value("aaa"))
}
