package client

import "net/http"

// SonrClient is a REST HTTP client for Querying Module Endpoints
// for the Sonr blockchain.
type SonrClient struct {
	apiURL string
}

// NewLocal creates a new SonrClient for local development.
func NewLocal() (*SonrClient, error) {
	// create http client
	client := &SonrClient{
		apiURL: "http://localhost:1323",
	}
	// Issue ping to check if server is up
	resp, err := http.Get(client.apiURL + "/genesis")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Check if server is up
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	return client, nil
}

// NewRemote creates a new SonrClient for remote production.
func NewRemote(url string) (*SonrClient, error) {
	// create http client
	client := &SonrClient{
		apiURL: url,
	}
	// Issue ping to check if server is up
	resp, err := http.Get(client.apiURL + "/genesis")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Check if server is up
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	return client, nil
}
