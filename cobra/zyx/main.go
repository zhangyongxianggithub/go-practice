package main

import (
	"fmt"

	"go-practice/cobra/zyx/cmd"
)

func main() {
	fmt.Println("command start executing")
	// main函数的唯一目的: 初始化 Cobra
	cmd.Execute()
	fmt.Println("command end executing")
}
