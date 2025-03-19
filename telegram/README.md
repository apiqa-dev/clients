# Telegram Client Library

A simple Go client library for sending messages via a Telegram server API.

## Features

- Send messages to Telegram channels via HTTP API
- Type-safe channel selection with custom Channel type
- Configurable timeout and retry options
- Simple and easy-to-use API
- Support for custom HTTP clients

## Installation

```bash
go get github.com/apiqa-dev/clients/telegram@latest
```

Or specify a specific version/commit:

```bash
go get github.com/apiqa-dev/clients/telegram@v1.0.0
```

## Usage

### Basic Usage

```go
package main

import (
	"log"
	
	"github.com/apiqa-dev/clients/telegram"
)

func main() {
	// Create a client with server URL and API key
	client := telegram.NewClient(
		"http://localhost:8080",     // Server URL
		"your-secure-api-key-here",  // API Key
	)

	// Send a message using predefined channel constants
	err := client.SendMessage(telegram.ChannelSugar, "Hello from the Telegram client library!")
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	
	// Send to another channel
	err = client.SendMessage(telegram.ChannelMBank, "This is a message to MBank channel!")
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
}
```

### Advanced Usage with Options

```go
package main

import (
	"log"
	"time"
	"net/http"
	
	"github.com/apiqa-dev/clients/telegram"
)

func main() {
	// Create a custom HTTP client
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 5,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	// Create a client with options
	client := telegram.NewClient(
		"https://api.example.com",   // Server URL
		"your-secure-api-key-here",  // API Key
		telegram.WithRetries(5),
		telegram.WithRetryDelay(2*time.Second),
		telegram.WithHTTPClient(httpClient),
	)

	// Send messages to different channels
	client.SendMessage(telegram.ChannelSugar, "Message to Sugar channel")
	client.SendMessage(telegram.ChannelLab, "Message to Lab channel")
	client.SendMessage(telegram.ChannelCommits, "New commit: Updated client library")
}
```

## Channel Type

The library uses a custom `Channel` type for type safety:

```go
type Channel string

const (
	ChannelSugar   Channel = "sugar"
	ChannelMBank   Channel = "mbank"
	ChannelLab     Channel = "lab"
	ChannelCommits Channel = "commits"
)
```

This provides compile-time safety - you can only pass a `Channel` type to the `SendMessage` method, not a regular string.

You can get a list of all available channels using:

```go
channels := telegram.AllChannels() // Returns []Channel{ChannelSugar, ChannelMBank, ChannelLab, ChannelCommits}
```

And check if a channel is valid using:

```go
isValid := telegram.IsValidChannel(telegram.ChannelSugar) // Returns true
isValid = telegram.IsValidChannel(telegram.Channel("invalid")) // Returns false
```

## Configuration Options

| Option | Description | Default |
|--------|-------------|---------|
| WithTimeout | Timeout for HTTP requests | 10 seconds |
| WithRetries | Number of retries for failed requests | 3 |
| WithRetryDelay | Delay between retries | 1 second |
| WithHTTPClient | Custom HTTP client | Default Go HTTP client |

## API Endpoints

The client expects the server to have the following endpoint:

- `POST /send` - Send a message to a channel

### Request Format

```json
{
  "channel": "channel_name",
  "message": "Your message text"
}
```

### Headers

- `Content-Type: application/json`
- `X-API-Key: your-api-key`

## Troubleshooting

### Common Issues

1. **Version Issues**: If you're getting version conflicts, try clearing your module cache:
   ```bash
   go clean -modcache
   ```

2. **Build Issues**: Make sure you're using the correct import path:
   ```go
   import "github.com/apiqa-dev/clients/telegram"
   ```

3. **Type Errors**: If you get a compile error about incompatible types, make sure you're using the `Channel` type:
   ```go
   // Correct
   client.SendMessage(telegram.ChannelSugar, "message")
   
   // Incorrect - will not compile
   client.SendMessage("sugar", "message")
   ```

## License

MIT 