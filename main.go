package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// StandupData struct to hold standup details
type StandupData struct {
	Date          string `json:"date"`
	YesterdayWork string `json:"yesterday_work"`
	TodayPlan     string `json:"today_plan"`
	Blockers      string `json:"blockers"`
}

func captureInput() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What did you work on yesterday? ")
	yesterdayWork, _ := reader.ReadString('\n')

	fmt.Print("What are you planning to work on today? ")
	todayPlan, _ := reader.ReadString('\n')

	fmt.Print("Any blocker?")
	blockers, _ := reader.ReadString('\n')

	// get the current date
	currentDate := time.Now().Format("2006-01-02")

	// create a new standup data entry
	standup := StandupData{
		Date:          currentDate,
		YesterdayWork: yesterdayWork,
		TodayPlan:     todayPlan,
		Blockers:      blockers,
	}

	// store the data in a local JSON file
	storeData(standup)
}

func storeData(data StandupData) {
	file, err := os.OpenFile("standup_data.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	// convert the standup data to JSON
	dataJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Failed to marshal data: %v\n", err)
		return
	}

	// write JSON data to the file
	if _, err := file.Write(append(dataJSON, '\n')); err != nil {
		fmt.Printf("Failed to write data: %v\n", err)
		return
	}

	fmt.Println("Standup data saved successfully!")
}

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

	// capture user input
	captureInput()
}

func main() {
	// Schedule standup bot to run at 9:00 Am every weekday
	scheduleStandup(9, 0)
}
