package httputil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FetchAndDecode makes a GET request to the specified URL and decodes the JSON response into the provided type T
func FetchAndDecode[T any](url string) (*T, error) {
	// Create HTTP client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Decode JSON into generic type
	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return &result, nil
}
