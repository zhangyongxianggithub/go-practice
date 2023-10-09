package main

import (
	"fmt"
)

func coinChangeGreedy(coins []int, amt int) int {
	// 假设coins按照大小有序，按照最大的金额先取
	i := len(coins) - 1
	count := 0
	for amt > 0 && i > 0 {
		for i > 0 && coins[i] > amt {
			i--
		}
		amt -= coins[i]
		count++
	}
	if amt != 0 {
		return -1
	}
	return count
}

func main() {
	coins := [7]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(coins[0])
}
