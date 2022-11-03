package main

import "fmt"

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func main3() {
	result := sumNumbers(1, 2, 3, 4)
	fmt.Println("the result is ", result)
}
