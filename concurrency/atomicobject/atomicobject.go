package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var config atomic.Value
	config.Store(loadConfig())
	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(loadConfig())
		}
	}()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for r := range requests() {
				time.Sleep(time.Second)
				c := config.Load()
				fmt.Println(c, i, r)
			}
		}(i)
	}
	wg.Wait()
}

func loadConfig() int64 {
	return time.Now().UnixNano()
}
func requests() []int {
	return []int{1, 2, 3}
}
