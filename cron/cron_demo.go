package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

/*
这是一个官方包内的定时调度库，支持cron表达式，支持固定周期，但是不考虑任务的执行时间，是一个比较简单的调度框架
*/
func main() {
	c := cron.New()
	c.AddFunc("0/1 * * * * *", func() { fmt.Println(time.Now(), "every second run") })
	c.AddFunc("0/10 * * * * *", func() { fmt.Println(time.Now(), "every ten seconds run") })

	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 10s", func() { fmt.Println(time.Now(), "Every ten seconds run @every") })
	c.Start()

	// Funcs are invoked in their own goroutine, asynchronously.
	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") })

	// Inspect the cron job entries' next and previous run times.
	inspect(c.Entries())
	// c.Stop() // Stop the scheduler (does not stop any jobs already running).
	time.Sleep(1 * time.Hour)
}

func inspect(entries []*cron.Entry) {
	for _, entry := range entries {
		fmt.Printf("job: %v, prev: %v, next: %v\n", entry.Job, entry.Prev, entry.Next)
	}
}
