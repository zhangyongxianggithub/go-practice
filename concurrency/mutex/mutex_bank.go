package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount

}
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	b := balance
	return b
}

func main() {
	go Deposit(10)
	go func() {
		fmt.Println(Balance())
	}()
	go Deposit(100)
	go func() {
		fmt.Println(Balance())
	}()
	time.Sleep(1 * time.Minute)
}
