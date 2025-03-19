package main

import (
	"fmt"
	"log"
	"time"

	"github.com/apiqa-dev/clients/telegram"
)

func main() {
	// Create a client with server URL and API key
	client := telegram.NewClient(
		"http://localhost:8080",    // Server URL
		"your-secure-api-key-here", // API Key
		telegram.WithRetries(5),
		telegram.WithRetryDelay(2*time.Second),
		telegram.WithTimeout(15*time.Second),
	)

	// Send messages to different channels using the predefined channel constants
	err := client.SendMessage(telegram.ChannelSugar, "Hello from the Sugar channel!")
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Println("Message sent to Sugar channel successfully!")

	err = client.SendMessage(telegram.ChannelMBank, "This is a message to MBank channel!")
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Println("Message sent to MBank channel successfully!")

	err = client.SendMessage(telegram.ChannelLab, "This is a message to Lab channel!")
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Println("Message sent to Lab channel successfully!")

	err = client.SendMessage(telegram.ChannelCommits, "New commit: Updated client library")
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Println("Message sent to Commits channel successfully!")

	// This would fail with a compile-time error because "invalid-channel" is not a Channel type
	// err = client.SendMessage("invalid-channel", "This message will not be sent")

	// This would fail with a runtime error because it's not a valid predefined channel
	invalidChannel := telegram.Channel("invalid-channel")
	err = client.SendMessage(invalidChannel, "This message will not be sent")
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// Print all available channels
	fmt.Println("Available channels:", telegram.AllChannels())
}
