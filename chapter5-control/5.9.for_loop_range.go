package main

import "fmt"

func main9() {
	numbers := []int{1, 2, 3, 4}
	for i, n := range numbers {
		fmt.Println("the index of the loop is ", i)
		fmt.Println("the value from the array is ", n)
	}
}
