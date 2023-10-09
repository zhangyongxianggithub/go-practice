package main

import (
	"sync"
	"sync/atomic"
)

type Singleton struct{}

var (
	instance    *Singleton
	initialized uint32
	mu          sync.Mutex // guards
)

func Instance() *Singleton {
	// 锁的代价很高，使用一个标志位避免过多的使用互斥锁
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	// 这个就是sync.Once的实现
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &Singleton{}
	}
	return instance
}
