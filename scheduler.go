package main

import (
	"fmt"
	"time"
)

// StartScheduler inits and starts scheduler
func StartScheduler(config Config) {
	go func() {
		for {
			now := time.Now()
			nextRun := time.Date(now.Year(), now.Month(), now.Day(), config.Hour, config.Minute, 0, 0, now.Location())

			// to test manually
			// nextRun := now.Add(2 * time.Minute)

			if now.After(nextRun) {
				nextRun = nextRun.Add(24 * time.Hour)
			}

			// Skip weekends
			for nextRun.Weekday() == time.Saturday || nextRun.Weekday() == time.Sunday {
				nextRun = nextRun.Add(24 * time.Hour)
			}

			waitDuration := nextRun.Sub(now)
			fmt.Printf("Next standup scheduled at: %v\n", nextRun.Format(time.RFC1123))

			time.Sleep(waitDuration)
			SendNotification()
		}
	}()
}
