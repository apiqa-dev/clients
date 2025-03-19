package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client represents a client for the Telegram message sending service
type Client struct {
	serverURL  string
	apiKey     string
	httpClient *http.Client
	timeout    time.Duration
	retries    int
	retryDelay time.Duration
}

// Config contains the configuration for the client
type Config struct {
	ServerURL string
	APIKey    string
}

// MessageRequest represents the request body for sending a message
type MessageRequest struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

// MessageResponse represents the response from the server
type MessageResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

// NewClient creates a new client instance
func NewClient(serverURL, apiKey string, opts ...Option) *Client {
	client := &Client{
		serverURL:  serverURL,
		apiKey:     apiKey,
		httpClient: &http.Client{},
		timeout:    10 * time.Second, // Default timeout
		retries:    3,                // Default retries
		retryDelay: time.Second,      // Default retry delay
	}

	// Apply options
	for _, opt := range opts {
		opt(client)
	}

	// Set timeout for HTTP client
	client.httpClient.Timeout = client.timeout

	return client
}

// SendMessage sends a message to a specified channel
func (c *Client) SendMessage(channel Channel, message string) error {
	// Validate that the channel is one of the predefined channels
	if !IsValidChannel(channel) {
		return fmt.Errorf("invalid channel: %s. Must be one of: %v", channel, AllChannels())
	}

	reqBody := MessageRequest{
		Channel: channel.String(),
		Message: message,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	// Implement retries
	var lastErr error
	for attempt := 0; attempt <= c.retries; attempt++ {
		if attempt > 0 {
			time.Sleep(c.retryDelay)
		}

		err = c.doRequest(jsonData)
		if err == nil {
			return nil
		}
		lastErr = err
	}

	return fmt.Errorf("failed to send message after %d attempts: %w", c.retries, lastErr)
}

// doRequest performs the actual HTTP request to the server
func (c *Client) doRequest(jsonData []byte) error {
	// Create request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/send", c.serverURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		var errResp MessageResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("server returned status %d", resp.StatusCode)
		}
		return fmt.Errorf("server returned error: %s", errResp.Error)
	}

	return nil
}
