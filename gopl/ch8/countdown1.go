package main

import (
	"fmt"
	"os"
	"time"
)

func main1() {
	fmt.Println("Commencing countdown")
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
		fmt.Printf("time channel: %s\n", <-tick)

	}
	launch()
}

func main() {
	// ...create abort channel...
	abort := make(chan struct{})
	go func() {
		_, err := os.Stdin.Read(make([]byte, 1))
		if err != nil {
			return
		} // read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		launch()
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

// func launch() {
// 	fmt.Println("launch")
// }
