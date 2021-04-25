package main

import "fmt"

func getPrize() (int, string) {
	i := 2
	s := "goldfish"
	return i, s
}
func main2() {
	quantity, prize := getPrize()
	fmt.Println("you own %v %v\n", quantity, prize)
}
