package main

import (
  "fmt"
  "time"
  "os/exec"
)

func scheduleStandup(hour, minute int) {
  // Get the current time
  now := time.Now()

  // Calculate next occurrence of the specified time
  nextRun := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())

  // If next run time is in the past, schedule it for the next day
  if now.After(nextRun) {
    nextRun = nextRun.Add(24 * time.Hour)
  }

  // Find the next weekday if today is a weekend
  for nextRun.Weekday() == time.Saturday || nextRun.Weekday() == time.Sunday {
    nextRun = nextRun.Add(24 * time.Hour)
  }

  // Calculate the duration to wati before running the standup bot
  waitDuration := nextRun.Sub(now)

  fmt.Printf("Standup bot scheduled to run in %v\n", waitDuration)

  // Wait for the specified duration
  time.Sleep(waitDuration)

  // Run the standup bot function
  runStandupBot()

  // Schedule the next standup after the current one completes
  go scheduleStandup(hour, minute)
}

func runStandupBot() {
  // Test command to send a notifcation on mac
  cmd := exec.Command("osascript", "-e", `display notification "Time for standup!" with title "Standup Bot"`)

  err := cmd.Run()
  if err != nil {
    fmt.Printf("Failed to send notification: %v\n", err)
  } else {
    fmt.Printf("Notification sent successfully")
  }

  // TODO: Capture user input
  // captureInput()
}


func main() {
  // Schedule standup bot to run at 9:00 Am every weekday
  scheduleStandup(9, 0)
}
