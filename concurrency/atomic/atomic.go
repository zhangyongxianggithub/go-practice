package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total uint64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64
	for ; i < 100; i++ {
		// total += i
		atomic.AddUint64(&total, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go worker(&wg)
	go worker(&wg)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total)
}
