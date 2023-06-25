package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("f.json")
	if err != nil {
		return
	}
	fmt.Println(string(file))
}
