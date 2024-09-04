package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config struct to hold config settings
type Config struct {
	Hour        int    `json:"hour"`
	Minute      int    `json:"minute"`
	StoragePath string `json:"storage_path"`
}

// LoadConfig loads config from JSON file
func LoadConfig() Config {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Config file not found. Using default settings.")
		return Config{
			Hour:        9,
			Minute:      0,
			StoragePath: "standup_data.json",
		}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding config file, using default settings.")
		return Config{
			Hour:        9,
			Minute:      0,
			StoragePath: "standup_data.json",
		}
	}
	return config
}
