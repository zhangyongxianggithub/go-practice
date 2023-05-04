package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.  Press return to abort.")
	abort := make(chan struct{})
	go func() {
		_, err := os.Stdin.Read(make([]byte, 1))
		if err != nil {
			return
		} // read a single byte
		abort <- struct{}{}
	}()
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}
func launch() {
	fmt.Println("launch")
}
