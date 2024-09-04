package main

import (
	"fmt"
	"os/exec"
)

// SendNotification sends system notification that launches input form when clicked
func SendNotification() {
	cmdNotification := exec.Command("osascript", "-e", `display notification "Time for your daily standup." with title "Standup Bot" sound name "default"`)
	err := cmdNotification.Run()
	if err != nil {
		fmt.Printf("Failed to send notification: %v\n", err)
	}

	OpenInputForm()
}
