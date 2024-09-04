package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// StandupData struct to hold standup details
type StandupData struct {
	Date          string `json:"date"`
	YesterdayWork string `json:"yesterday_work"`
	TodayPlan     string `json:"today_plan"`
	Blockers      string `json:"blockers"`
}

// StoreData saves standup info to JSON file
func StoreData(data StandupData) error {
	filePath := LoadConfig().StoragePath

	var standups []StandupData

	// Check if file exists
	if _, err := os.Stat(filePath); err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("Failed to open file: %v\n", err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&standups)
		if err != nil {
			return fmt.Errorf("Failed to decode existing data: %v\n", err)
		}
	}

	standups = append(standups, data)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Failed to create file: %v\n", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(standups)
	if err != nil {
		return fmt.Errorf("Failed to encode data: %v\n", err)
	}

	return nil
}
