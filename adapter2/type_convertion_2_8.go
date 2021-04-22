package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var s = "true"
	b, err := strconv.ParseBool(s)
	fmt.Println(b, reflect.TypeOf(b), err)
	s = strconv.FormatBool(b)
	fmt.Println(s)
}
