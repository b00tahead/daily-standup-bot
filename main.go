package main

func main() {
	// Load config settings
	config := LoadConfig()

	// Start scheduler
	StartScheduler(config)
}
