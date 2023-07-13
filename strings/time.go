package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(fmt.Sprintf("%s", time.Now().Format("20060102150405.000000")))
}
