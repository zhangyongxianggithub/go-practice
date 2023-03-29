package main

import (
	"fmt"
	"time"
)

func main() {
	names := [...]string{"1", "2", "3", "4"}
	ch := make(chan Result, len(names))
	for i := 0; i < len(names); i++ {
		go getResource(names[i], ch)
	}
	for range names {
		value := <-ch
		fmt.Println(value)
	}
}

func getResource(name string, ch chan Result) {
	// time.Sleep(1 * time.Second)
	defer func() {
		if p := recover(); p != nil {
			ch <- Result{name, name + time.Now().String(), fmt.Errorf("%v", p)}
		}
	}()
	fmt.Println("execute")
	if name == "2" {

		// panic("2 panic")
		time.Sleep(5 * time.Second)
	}
	ch <- Result{name, name + time.Now().String(), nil}
}

type Result struct {
	ant    interface{}
	result interface{}
	err    error
}
