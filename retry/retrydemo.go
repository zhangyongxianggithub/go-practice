package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sethvargo/go-retry"
)

func main() {
	ctx := context.Background()
	count := 0
	backOff := retry.NewFibonacci(1 * time.Second)
	backOff = retry.WithJitter(1*time.Second, backOff)
	backOff = retry.WithMaxRetries(2, backOff)
	if err := retry.Do(ctx, backOff, func(ctx context.Context) error {
		fmt.Printf("%s: %dth execute\n", time.Now(), count)
		count += 1
		return retry.RetryableError(fmt.Errorf("test"))
	}); err != nil {
		fmt.Println(err)
	}
}
